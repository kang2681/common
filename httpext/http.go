package httpext

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

const (
	OPTIONS string = "OPTIONS"
	GET     string = "GET"
	HEAD    string = "HEAD"
	POST    string = "POST"
	PUT     string = "PUT"
	DELETE  string = "DELETE"
	TRACE   string = "TRACE"
	CONNECT string = "CONNECT"
)

var DefaultTransport = &http.Transport{
	DisableKeepAlives: true,
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

type Client struct {
	C *http.Client
}

func (c *Client) GetHttpClient() *http.Client {
	return c.C
}

func NewClient() *Client {
	return NewClientWithContext(context.Background())
}
func NewClientWithContext(ctx context.Context) *Client {
	return &Client{
		C: &http.Client{
			Transport: DefaultTransport,
		},
	}
}

func (c *Client) SetTransport(transport http.RoundTripper) *Client {
	if c.C != nil {
		c.C.Transport = transport
	}
	return c
}

func (c *Client) SetCookieJar(jar http.CookieJar) *Client {
	if c.C != nil {
		c.C.Jar = jar
	}
	return c
}

func (c *Client) SetTimout(timeout time.Duration) *Client {
	if c.C != nil {
		c.C.Timeout = timeout
	}
	return c
}

type Request struct {
	Method      string         // 请求方法
	Url         string         // 请求地址
	Query       url.Values     // 地址栏参数
	Form        url.Values     // HTTP body 只有为nil 时，在写写入body 中，body不为空，直接丢弃
	Body        io.Reader      // http body
	ContentType string         // contentType default application/x-www-form-urlencoded
	Header      http.Header    // http header
	Cookies     []*http.Cookie // http cookie
}

type Response struct {
	r *http.Response
}

func ValidMethod(method string) bool {
	return method == OPTIONS || method == GET || method == HEAD || method == POST || method == PUT || method == DELETE || method == TRACE || method == CONNECT
}

func (c *Client) Do(ctx context.Context, p Request) (*Response, error) {
	p, err := c.check(p)
	if err != nil {
		return nil, fmt.Errorf("request param check error %s", err.Error())
	}
	r, err := http.NewRequestWithContext(ctx, p.Method, p.Url, p.Body)
	if err != nil {
		return nil, fmt.Errorf("new request error %s", err.Error())
	}
	// 设置HTTP header
	r.Header = p.Header
	if p.ContentType != "" {
		r.Header.Set("Content-Type", p.ContentType)
	}
	for _, v := range p.Cookies {
		r.AddCookie(v)
	}
	resp, err := c.C.Do(r)
	if err != nil {
		return nil, fmt.Errorf("send request error %s", err.Error())
	}
	return &Response{r: resp}, nil
}

func (c *Client) check(p Request) (Request, error) {
	if p.Method == "" {
		p.Method = GET
	}
	if p.Url == "" {
		return p, fmt.Errorf("url empty")
	}
	if p.ContentType == "" {
		p.ContentType = "application/x-www-form-urlencoded"
	}
	if p.Form != nil && len(p.Form) > 0 && p.Body == nil {
		p.Body = strings.NewReader(p.Form.Encode())
	}
	if p.Query != nil && len(p.Query) > 0 {
		if strings.Contains(p.Url, "?") {
			p.Url = p.Url + "&" + p.Query.Encode()
		} else {
			p.Url = p.Url + "?" + p.Query.Encode()
		}
	}
	return p, nil
}

func (c *Client) PostForm(url string, data url.Values) {
}

func (c *Client) Post(url string, data url.Values, body io.Reader) {
}

func (r *Response) Parse() ([]byte, error) {
	defer r.r.Body.Close()
	if r.r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response code %s not %d", r.r.StatusCode, http.StatusOK)
	}
	data, err := ioutil.ReadAll(r.r.Body)
	if err != nil {
		return nil, fmt.Errorf("read reponse body error %w", err)
	}
	return data, nil
}

func (r *Response) ParseWithStatus() ([]byte, int, error) {
	defer r.r.Body.Close()
	data, err := ioutil.ReadAll(r.r.Body)
	if err != nil {
		return nil, http.StatusOK, fmt.Errorf("read reponse body error %w", err)
	}
	return data, http.StatusOK, nil
}

type ChunkedData struct {
	Data []byte
	Err  error
}

func (r *Response) ParseChunked(dataChan chan<- ChunkedData) {
	if r.r.StatusCode != http.StatusOK {
		dataChan <- ChunkedData{Err: fmt.Errorf("response code %s not %d", r.r.StatusCode, http.StatusOK)}
		return
	}
	defer r.r.Body.Close()
	reader := httputil.NewChunkedReader(r.r.Body)
	for {
		reader.Read()
	}
}
