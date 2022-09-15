package khttp

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	resty.Client
}
