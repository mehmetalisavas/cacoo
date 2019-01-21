package cacoo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestClient tests the client and its default values
func TestClient(t *testing.T) {
	c := NewClient("")
	if c == nil {
		t.Error("client should not be nil")
	}
	if c.BaseURL.String() == "" {
		t.Error("client base url should not be empty")
	}
	if c.apiKey != "" {
		t.Error("client api key should be empty")
	}
	if c.UserAgent == "" {
		t.Error("client user-agent should not be empty")
	}
}

func TestClientOptionBaseURL(t *testing.T) {
	u := "https://test.com"
	c := NewClient("", OptionBaseURL(u))
	if c == nil {
		t.Error("client should not be nil")
	}
	if c.BaseURL.String() == "" {
		t.Error("client base url should not be empty")
	}
	if c.BaseURL.String() != u {
		t.Errorf("client base url should equal to: %s, but got: %s", u, c.BaseURL.String())
	}
}

func TestClientOptionUserAgent(t *testing.T) {
	u := "test-user-agent"
	c := NewClient("", OptionUserAgent(u))
	if c == nil {
		t.Error("client should not be nil")
	}
	if c.UserAgent != u {
		t.Errorf("client user agent should equal to: %s, but got: %s", c.UserAgent, u)
	}
}
func TestClientToken(t *testing.T) {
	token := "test12345"
	c := NewClient(token)
	if c == nil {
		t.Error("client should not be nil")
	}
	if c.apiKey != token {
		t.Errorf("client api key should equal to: %s, but got: %s", c.apiKey, token)
	}
}

const (
	baseTestURL = "/api/"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	rawurl, _ := url.Parse(server.URL)
	client = NewClient("")
	client.BaseURL = rawurl

	return client, mux, server.URL, server.Close
}

func method(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

// Bool returns a pointer to given value.
func Bool(v bool) *bool { return &v }

// Int returns a pointer to given value.
func Int(v int) *int { return &v }

// Int64 returns a pointer to given value.
func Int64(v int64) *int64 { return &v }

// String returns a pointer to given value.
func String(v string) *string { return &v }
