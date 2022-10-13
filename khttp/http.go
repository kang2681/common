package khttp

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	resty.Client
}

func Get(url string) (*resty.Response, error) {
	return resty.New().NewRequest().Get(url)
}

func Post(url string, data map[string]string) (*resty.Response, error) {
	return resty.New().NewRequest().SetFormData(data).Post(url)
}

func PostRaw(url string, body string) (*resty.Response, error) {
	return resty.New().NewRequest().SetBody(body).Post(url)
}
