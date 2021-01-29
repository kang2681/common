package httpext

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var DefaultTransport = &http.Transport{
	DisableKeepAlives: true,
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

type Client struct {
	c     *http.Client
	Query url.Values // 地址栏参数
}

type Config struct {
	Timeout time.Duration // 超时时间
}

func NewClient(opts ...ClientOptionFunc) (*Client, error) {
	return DialContext(context.Background(), opts...)
}

func DialContext(ctx context.Context, opts ...ClientOptionFunc) (*Client, error) {
	c := &Client{}
	// Run the options on it
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Client) Do() {
}

func (c *Client) Get() {

}

func (c *Client) Post() {

}

func (c *Client) Raw() {

}

type ClientOptionFunc func(*Client) error

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

func SetRoundTripper(transport http.RoundTripper) ClientOptionFunc {
	return func(c *Client) error {
		if c.c != nil {
			return fmt.Errorf("http client nil")
		}
		c.c.Transport = transport
		return nil
	}
}

func SetTransport(transport *http.Transport) ClientOptionFunc {
	return func(c *Client) error {
		if c.c != nil {
			return fmt.Errorf("http client nil")
		}
		c.c.Transport = transport
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
