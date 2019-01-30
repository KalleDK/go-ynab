package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/kalledk/go-ynab/ynab"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const ApiEndpoint = "https://api.youneedabudget.com/v1/"

type Client struct {
	Endpoint    *url.URL
	AccessToken ynab.AccessToken
	Client      HttpClient
	/*
		sync.Mutex
		rateLimit *api.RateLimit
	*/
}

func (c *Client) Get(path string, responseModel interface{}) (err error) {
	return c.Do(http.MethodGet, path, responseModel, nil)
}

func (c *Client) Do(method string, path string, responseModel interface{}, reqbody []byte) (err error) {
	fmt.Println(path)
	requrl, err := c.Endpoint.Parse(path)
	if err != nil {
		return err
	}
	fmt.Println(requrl)
	req, err := http.NewRequest(method, requrl.String(), bytes.NewBuffer(reqbody))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		var errorResponse ynab.ErrorResponse
		if err = json.Unmarshal(body, &errorResponse); err != nil {
			return err
		}

		return errorResponse.Error
	}

	//fmt.Println(string(body[:]))

	return json.Unmarshal(body, &responseModel)
}

func NewClient(accessToken ynab.AccessToken) (client *Client) {

	endpoint, err := url.Parse(ApiEndpoint)
	if err != nil {
		log.Fatalf("invalid endpoint %v", ApiEndpoint)
	}

	client = &Client{
		Endpoint:    endpoint,
		AccessToken: accessToken,
		Client:      http.DefaultClient,
	}
	return
}
