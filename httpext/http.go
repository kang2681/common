package httpext

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	c *http.Client
}

type Config struct {
	Timeout time.Duration // 超时时间
}

type ClientOptionFunc func(*Client) error

func NewClient(opts ...ClientOptionFunc) (*Client, error) {
	return DialContext(context.Background(), opts...)
}

func DialContext(ctx context.Context, opts ...ClientOptionFunc) (*Client, error) {
	return &Client{}, nil
}

func (c *Client) Do() {
}

func (c *Client) Get() {

}

func (c *Client) Post() {

}

func (c *Client) Raw() {

}

// SetTimeout 设置超时时间
func SetTimeout(timeout time.Duration) ClientOptionFunc {
	return func(c *Client) error {
		if c.c != nil {
			return fmt.Errorf("http client nil")
		}
		c.c.Timeout = timeout
		return nil
	}
}

// SetHttpClient 设置 httpClient
func SetHttpClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		if httpClient != nil {
			c.c = httpClient
		} else {
			c.c = http.DefaultClient
		}
		return nil
	}
}
