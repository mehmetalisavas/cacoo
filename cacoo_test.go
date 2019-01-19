package cacoo

import (
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
