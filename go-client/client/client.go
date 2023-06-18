package client

import "net/http"

var (
	DefaultAPIURL = "https://pokeapi.co/"
)

type Client struct {
	apiUrl     string
	httpClient *http.Client
}

type Option func(c *Client)

func NewClient(opts ...Option) *Client {
	client := &Client{
		apiUrl:     DefaultAPIURL,
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func WithAPIURL(url string) Option {
	return func(c *Client) {
		c.apiUrl = url
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
