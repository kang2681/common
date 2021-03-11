package httpext

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Method      string      // 请求方法
	Url         string      // URL 地址
	Query       url.Values  // 地址栏参数
	Form        url.Values  // 表单参数
	Body        io.Reader   // body 内容
	ContentType string      // Content-Type
	Header      http.Header // http header
	Cookies     []*http.Cookie
	Charset     string
}

func NewRequest(method string) *Request {
	return &Request{
		Method: method,
	}
}

func (r *Request) SetMethod(method string) *Request {
	r.Method = method
	return r
}

func (r *Request) SetURL(url string) *Request {
	r.Url = url
	return r
}

func (r *Request) SetForm(data url.Values) *Request {
	r.Form = data
	return r
}

func (r *Request) SetBody(body []byte) *Request {
	r.Body = bytes.NewReader(body)
	return r
}

func (r *Request) SetHeader(header http.Header) *Request {
	r.Header = header
	return r
}

func (r *Request) AddHeader(key, value string) *Request {
	r.Header.Set(key, value)
	return r
}

func (r *Request) AddHeaderForm(key, value string) *Request {
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r
}

func (r *Request) SetCharset(charset string) *Request {
	r.Charset = charset
	return r
}

func (r *Request) SetCookies(cookies []*http.Cookie) *Request {
	r.Cookies = cookies
	return r
}

func (r *Request) AddCookie(cookie *http.Cookie) *Request {
	r.Cookies = append(r.Cookies, cookie)
	return r
}
