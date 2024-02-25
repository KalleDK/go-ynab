package ynab

import "net/http"

// Scheme is what scheme is used to talkt to the api
const Scheme = "https://"

// BaseURL is the default domain to use
const BaseURL = "api.youneedabudget.com/v1"

// HTTPClient is the interface that is needed to reach the internet
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Config is how you want the client to work
type Config struct {
	Token   Token
	Scheme  string
	BaseURL string
	Client  HTTPClient
}

// NewClient creates a new client based of the config
func (c Config) NewClient() *Client {
	baseURL := BaseURL
	if c.BaseURL != "" {
		baseURL = c.BaseURL
	}

	scheme := Scheme
	if c.Scheme != "" {
		scheme = c.Scheme
	}

	var client HTTPClient = http.DefaultClient
	if c.Client != nil {
		client = c.Client
	}

	return &Client{
		token:   c.Token,
		scheme:  scheme,
		baseURL: baseURL,
		client:  client,
	}
}
