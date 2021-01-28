package httpext

import (
	"net/http"
	"time"
)

type Client struct {
	c *http.Client
}

type ClientOptionFunc func(*Client)

func SetTimeout(timeout time.Duration) ClientOptionFunc {
	return func(c *Client) {
		c.c.Timeout = timeout
	}
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Do() {

}

func (c *Client) Get() {

}

func (c *Client) Post() {

}

func (c *Client) Raw() {

}
