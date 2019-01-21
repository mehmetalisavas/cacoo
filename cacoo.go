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
func OptionHttpClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		c.client = httpClient
	}
}

// OptionUserAgent provides a user agent description for the cacoo client
func OptionUserAgent(agent string) func(*Client) {
	return func(c *Client) {
		c.UserAgent = agent
	}
}

// OptionBaseURL provides a specific url for the client. It's better to be
// careful while using this option. It's also useful for writing tests for the
// client, so you can modify your client endpoint
func OptionBaseURL(rawurl string) func(*Client) {
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

type Response struct {
	*http.Response
}

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

func (c *Client) Get(ctx context.Context, url string, v interface{}) (*Response, error) {
	return c.send(ctx, "GET", url, nil, v)
}
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
