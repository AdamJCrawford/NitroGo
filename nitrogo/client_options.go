package nitrogo

import "net/http"

type ClientOptionFunc func(*Client) error

func WithBaseUrl(url string) ClientOptionFunc {
	return func(c *Client) error {
		c.baseURL = url
		return nil
	}
}

func WithProxiedURL(url string) ClientOptionFunc {
	return func(c *Client) error {
		c.proxiedURL = url
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}
