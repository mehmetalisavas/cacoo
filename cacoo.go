package cacoo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://cacoo.com/"
	apiVersion     = "api/v1/"
	userAgent      = "cacoo-go-client"
)

// Client represents the Cacoo client information and services
type Client struct {
	// Client is used to communicate with the cacoo api
	client *http.Client

	apiKey string

	// BaseURL holds the url for the cacoo api
	BaseURL *url.URL

	// UserAgent holds the agent name while communicating with cacoo api
	UserAgent string

	// service is used to register all other services
	service service

	Users   *UsersService
	Account *AccountService
	Folders *FolderService
	Diagram *DiagramService
	Image   *ImageService
	License *LicenseService
	Comment *CommentService
}

// service holds the cacoo client and its services
type service struct {
	client *Client
}

// Option is used for client options
type Option func(*Client)

// NewClient returns the cacoo api client. Please provide a token to get the
// response which require authentication
//
// Second parameter(Options) is used for options of the client initialization.
// You can specifiy custom settings while creating the client
//
// you can initialize client basically;
// e.g;
// NewClient("") without token&option
// NewClient("apiKeyToken") with token
func NewClient(token string, opts ...Option) *Client {
	client := http.DefaultClient

	u := defaultBaseURL + apiVersion
	baseURL, _ := url.Parse(u)

	c := &Client{client: client, BaseURL: baseURL, UserAgent: userAgent, apiKey: token}
	c.service.client = c

	c.Users = (*UsersService)(&c.service)
	c.Account = (*AccountService)(&c.service)
	c.Folders = (*FolderService)(&c.service)
	c.Diagram = (*DiagramService)(&c.service)
	c.Image = (*ImageService)(&c.service)
	c.License = (*LicenseService)(&c.service)
	c.Comment = (*CommentService)(&c.service)

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// OptionHttpClient provides a client for cacoo cliet
// usage is;
// NewClient("apikey", OptionHttpClient(yourHttpClient *http.Client))
func OptionHttpClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.client = httpClient
	}
}

// OptionUserAgent provides a user agent description for the cacoo client
// simple usage is; // NewClient("apikey", OptionUserAgent("custom user agent"))
func OptionUserAgent(agent string) Option {
	return func(c *Client) {
		c.UserAgent = agent
	}
}

// OptionBaseURL provides a specific url for the client. It's better to be
// careful while using this option. It's also useful for writing tests for the
// client, so you can modify your client endpoint
func OptionBaseURL(rawurl string) Option {
	return func(c *Client) {
		u, _ := url.Parse(rawurl)
		c.BaseURL = u
	}
}

// NewRequest creates the cacoo api request
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		encode := json.NewEncoder(buf)
		encode.SetEscapeHTML(false)
		err = encode.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if c.apiKey != "" {
		q := req.URL.Query()
		q.Add("apiKey", c.apiKey)
		req.URL.RawQuery = q.Encode()
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Response represents the response struct of cacoo client
type Response struct {
	*http.Response
}

// ErrorResponse represent the returning error struct.
// Implements the error interface
type ErrorResponse struct {
	Response *http.Response
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Content  string `json:"content"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message, r.Content)
}

// Get used for GET methods of cacoo client
func (c *Client) Get(ctx context.Context, url string, v interface{}) (*Response, error) {
	return c.send(ctx, "GET", url, nil, v)
}

// Post used for POST methods of cacoo client
func (c *Client) Post(ctx context.Context, url string, body, v interface{}) (*Response, error) {
	return c.send(ctx, "POST", url, body, v)
}

func (c *Client) send(ctx context.Context, method string, url string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, v)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do sends the requests and returns the cacoo api response.  If there is an
// error while doing request then it returns error, and sanitizes the url if
// necessary  if v implements the writer interface then it copies the body to
// the given parameter v. Otwerwise just decodes the incoming data into the v
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if e, ok := err.(*url.Error); ok {
			if u, err := url.Parse(e.URL); err == nil {
				// sanitize the url if error is *url.Error
				e.URL = sanitizeURL(u).String()
				return nil, e
			}
		}
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

// sanitizeURL redacts the apiKey parameter from the URL
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("apiKey")) > 0 {
		params.Set("apiKey", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

// CheckResponse checks the status code of the response, if there is an error it
// parses incoming data (if it exists)
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
