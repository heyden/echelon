package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

var defaultHttpClient = &Client{client: http.DefaultClient}

type Client struct {
	client      *http.Client
	baseURL     string
	baseHeaders http.Header
}

type ClientOptions struct {
	HttpClient *http.Client
	baseURL    string
	Headers    http.Header
}

func NewClient(o ClientOptions) *Client {
	c := &Client{
		baseURL:     o.baseURL,
		baseHeaders: o.Headers,
	}

	if o.HttpClient != nil {
		c.client = o.HttpClient
	} else {
		c.client = http.DefaultClient
	}

	return c
}

func (c *Client) Get(ctx context.Context, url string, headers http.Header, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req, v)
}

func (c *Client) Post(ctx context.Context, url string, headers http.Header, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req, v)
}

func (c *Client) Put(ctx context.Context, url string, headers http.Header, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req, v)
}

func (c *Client) Patch(ctx context.Context, url string, headers http.Header, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req, v)
}

func (c *Client) Delete(ctx context.Context, url string, headers http.Header, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req, v)
}

// do sends a request and returns the response. It can decode JSON responses in the value pointed to by v.
func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)

	if v != nil {
		defer resp.Body.Close()
		decodeErr := json.NewDecoder(resp.Body).Decode(v)
		if decodeErr == io.EOF {
			decodeErr = nil
		}
		if decodeErr != nil {
			err = decodeErr
		}
	}
	return resp, err
}

func Get(ctx context.Context, uri string, headers http.Header, v interface{}) (*http.Response, error) {
	return defaultHttpClient.Get(ctx, uri, headers, v)
}

func Post(ctx context.Context, uri string, headers http.Header, v interface{}) (*http.Response, error) {
	return defaultHttpClient.Post(ctx, uri, headers, v)
}

func Put(ctx context.Context, uri string, headers http.Header, v interface{}) (*http.Response, error) {
	return defaultHttpClient.Put(ctx, uri, headers, v)
}

func Patch(ctx context.Context, uri string, headers http.Header, v interface{}) (*http.Response, error) {
	return defaultHttpClient.Patch(ctx, uri, headers, v)
}

func Delete(ctx context.Context, uri string, headers http.Header, v interface{}) (*http.Response, error) {
	return defaultHttpClient.Delete(ctx, uri, headers, v)
}

func EncodeJson(v interface{}) (*bytes.Buffer, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}
