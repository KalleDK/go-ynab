package client

import "net/url"

type JSONClient interface {
	Get(reqURL *url.URL, responseModel interface{}) (err error)
	Post(reqURL *url.URL, data interface{}, responseModel interface{}) (err error)
}

type APIEndpoint struct {
	URL        *url.URL
	JSONClient JSONClient
}

func (a *APIEndpoint) GetVia(path string, responseModel interface{}) (err error) {
	reqURL, err := a.URL.Parse(path)
	if err != nil {
		return
	}
	return a.JSONClient.Get(reqURL, responseModel)
}

func (a *APIEndpoint) PostVia(path string, data interface{}, responseModel interface{}) (err error) {
	reqURL, err := a.URL.Parse(path)
	if err != nil {
		return
	}
	return a.JSONClient.Post(reqURL, data, responseModel)
}
