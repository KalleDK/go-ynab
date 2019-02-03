package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type AuthClient interface {
	AddAuth(header *http.Header)
}

type ReflectJSONClient struct {
	ErrorModel interface{}
	HTTPClient HTTPClient
}

func (c *ReflectJSONClient) Do(method string, reqURL *url.URL, responseModel interface{}, reqbody []byte) (err error) {

	fmt.Printf("%ving: %v\n", method, reqURL)

	req, err := http.NewRequest(method, reqURL.String(), bytes.NewBuffer(reqbody))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		errorModel := reflect.TypeOf(c.ErrorModel)
		errorPtr := reflect.New(errorModel)

		if err = json.Unmarshal(body, errorPtr.Interface()); err != nil {
			return err
		}

		return errorPtr.Elem().Interface().(error)
	}

	return json.Unmarshal(body, &responseModel)
}

func (c *ReflectJSONClient) Get(reqURL *url.URL, responseModel interface{}) (err error) {
	return c.Do(http.MethodGet, reqURL, responseModel, nil)
}

func (c *ReflectJSONClient) Post(reqURL *url.URL, data interface{}, responseModel interface{}) (err error) {
	body, err := json.Marshal(data)
	if err != nil {
		return
	}

	return c.Do(http.MethodPost, reqURL, responseModel, body)
}
