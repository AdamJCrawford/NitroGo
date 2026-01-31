package nitrogo

import "net/http"

type ClientOptionFunc func(*Client) error

func WithHostname(hostname string) ClientOptionFunc {
	return func(c *Client) error {
		c.hostname = hostname
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}

func WithProxiedURL(url string) ClientOptionFunc {
	return func(c *Client) error {
		c.proxiedURL = url
		return nil
	}
}
