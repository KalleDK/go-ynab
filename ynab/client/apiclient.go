package client

import "net/url"

type JsonClient interface {
	Get(requrl *url.URL, responseModel interface{}) (err error)
	Post(requrl *url.URL, data interface{}, responseModel interface{}) (err error)
}

type APIEndpoint struct {
	Url        *url.URL
	JsonClient JsonClient
}

func (a *APIEndpoint) GetVia(path string, responseModel interface{}) (err error) {
	reqUrl, err := a.Url.Parse(path)
	if err != nil {
		return
	}
	return a.JsonClient.Get(reqUrl, responseModel)
}

func (a *APIEndpoint) PostVia(path string, data interface{}, responseModel interface{}) (err error) {
	reqUrl, err := a.Url.Parse(path)
	if err != nil {
		return
	}
	return a.JsonClient.Post(reqUrl, data, responseModel)
}
