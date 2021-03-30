package httpext

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	urlpkg "net/url"
	"strings"
	"time"
)

type ContentType = string

var (
	Content_Text_Plain     ContentType = "text/plain"
	Content_Text_Html      ContentType = "text/html"
	Content_Text_Csv       ContentType = "text/csv"
	Content_App_Json       ContentType = "application/json"
	Content_App_Javascript ContentType = "application/javascript"
	Content_App_Xml        ContentType = "application/xml"
	Content_App_Form       ContentType = "application/x-www-form-urlencoded"
)

type Client struct {
	c *http.Client
}

type ClientOption func(c *Client)

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		c: &http.Client{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func NewDefaultClient() *Client {
	c := &Client{
		c: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				DisableKeepAlives: true,
			},
		},
	}
	return c
}

func (c *Client) SetOptions(opts ...ClientOption) *Client {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func DialTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.c.Timeout = timeout
	}
}

func DialTransport(transport *http.Transport) ClientOption {
	return func(c *Client) {
		c.c.Transport = transport
	}
}

func DialRoundTripper(transport http.RoundTripper) ClientOption {
	return func(c *Client) {
		c.c.Transport = transport
	}
}

func DialCookieJar(jar http.CookieJar) ClientOption {
	return func(c *Client) {
		c.c.Jar = jar
	}
}

func DialCheckRedirect(fn func(req *http.Request, via []*http.Request) error) ClientOption {
	return func(c *Client) {
		c.c.CheckRedirect = fn
	}
}

type Request struct {
	URL     string         // 请求的URL
	Query   urlpkg.Values  // 请求地址栏参数
	Header  http.Header    // http 请求头
	Cookies []*http.Cookie // http cookie
}

type GetRequest struct {
	Request
	Body io.Reader
}

func (r *Request) URLAddr() string {
	buf := bytes.NewBufferString(r.URL)
	if len(r.Query) > 0 {
		if strings.Contains(r.URL, "?") {
			buf.WriteString("&")
		} else {
			buf.WriteString("?")
		}
		buf.WriteString(r.Query.Encode())
	}
	return buf.String()
}

func (c *Client) Get(ctx context.Context, r *GetRequest) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", r.URLAddr(), r.Body)
	if err != nil {
		return nil, err
	}
	if len(r.Header) > 0 {
		req.Header = r.Header.Clone()
	}
	for _, v := range r.Cookies {
		req.AddCookie(v)
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{
		Resp: resp,
	}, nil
}

type PostFormRequest struct {
	Request
	FormData urlpkg.Values
	Charset  string // charset default UTF-8
}

func (c *Client) PostForm(ctx context.Context, r *PostFormRequest) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", r.URLAddr(), strings.NewReader(r.FormData.Encode()))
	if err != nil {
		return nil, err
	}
	if len(r.Header) > 0 {
		req.Header = r.Header.Clone()
	}
	for _, v := range r.Cookies {
		req.AddCookie(v)
	}
	if val := req.Header.Values("Content-Type"); len(val) == 0 {
		var charset string
		if r.Charset == "" {
			charset = "utf-8"
		}
		req.Header.Set("Content-Type", fmt.Sprintf("%s; charset=%s", Content_App_Form, charset))
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{
		Resp: resp,
	}, nil
}

type PostRequest struct {
	Request
	Body        io.Reader
	ContentType string
}

// text/plain
func (c *Client) PostRaw(ctx context.Context, r *PostRequest) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", r.URLAddr(), r.Body)
	if err != nil {
		return nil, err
	}
	if len(r.Header) > 0 {
		req.Header = r.Header.Clone()
	}
	for _, v := range r.Cookies {
		req.AddCookie(v)
	}
	if val := req.Header.Values("Content-Type"); len(val) == 0 {
		var contentType string
		if contentType == "" {
			contentType = Content_Text_Plain
		}
		req.Header.Set("Content-Type", contentType)
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{
		Resp: resp,
	}, nil
}

type PostFileRequest struct {
	Request
	FormData urlpkg.Values
	FormFile []PostFormFile
}

type PostFormFile struct {
	Field    string
	Filename string
	Content  []byte
}

func (c *Client) PostFile(ctx context.Context, r *PostFileRequest) (*Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	if err := c.postFileWriteForm(bodyWriter, r.FormData, r.FormFile); err != nil {
		bodyWriter.Close()
		return nil, err
	}
	bodyWriter.Close()
	req, err := http.NewRequestWithContext(ctx, "POST", r.URLAddr(), bodyBuf)
	if err != nil {
		return nil, err
	}
	if len(r.Header) > 0 {
		req.Header = r.Header.Clone()
	}
	for _, v := range r.Cookies {
		req.AddCookie(v)
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{
		Resp: resp,
	}, nil
}

func (c *Client) postFileWriteForm(bodyWriter *multipart.Writer, formData urlpkg.Values, formFiles []PostFormFile) error {
	defer bodyWriter.Close()
	for k, v := range formData {
		for _, vv := range v {
			bodyWriter.WriteField(k, vv)
		}
	}
	for _, v := range formFiles {
		fileWriter, err := bodyWriter.CreateFormFile(v.Field, v.Filename)
		if err != nil {
			return err
		}
		if _, err := fileWriter.Write(v.Content); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) urlGen(url string, queryData urlpkg.Values) string {
	if len(queryData) <= 0 {
		return url
	}
	if strings.Contains(url, "?") {
		url = url + "&" + queryData.Encode()
	} else {
		url = url + "?" + queryData.Encode()
	}
	return url
}

type Response struct {
	Resp *http.Response
}

type RespChunkedData struct {
	Err  error
	Data []byte
}

func (r *Response) Code() int {
	return r.Resp.StatusCode
}

func (r *Response) Body() ([]byte, error) {
	return ioutil.ReadAll(r.Resp.Body)
}

func (r *Response) Close() error {
	return r.Resp.Body.Close()
}

func (r *Response) Data() ([]byte, error) {
	if r.Resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response code %d not %d", r.Resp.StatusCode, http.StatusOK)
	}
	data, err := ioutil.ReadAll(r.Resp.Body)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Response) Cookies() []*http.Cookie {
	return r.Resp.Cookies()
}

func (r *Response) Json(v interface{}) error {
	if r.Resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response code %d not %d", r.Resp.StatusCode, http.StatusOK)
	}
	data, err := ioutil.ReadAll(r.Resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return err
}

func (r *Response) Chunked() <-chan RespChunkedData {
	dataChan := make(chan RespChunkedData, 10)
	go r.doChunked(dataChan)
	return dataChan
}

func (r *Response) doChunked(dataChan chan RespChunkedData) {
	defer close(dataChan)
	reader := httputil.NewChunkedReader(r.Resp.Body)
	buf := make([]byte, 4096)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				dataChan <- RespChunkedData{Err: err}
			} else if n > 0 {
				data := make([]byte, 0, n)
				data = append(data, buf[0:n]...)
				dataChan <- RespChunkedData{Data: data}
			}
			return
		}
		if n > 0 {
			data := make([]byte, 0, n)
			data = append(data, buf[0:n]...)
			dataChan <- RespChunkedData{Data: data}
		}
	}
}
