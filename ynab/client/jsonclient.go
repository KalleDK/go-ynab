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

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type AuthClient interface {
	AddAuth(header *http.Header)
}

type ReflectJsonClient struct {
	ErrorModel interface{}
	HttpClient HttpClient
}

func (c *ReflectJsonClient) Do(method string, regUrl *url.URL, responseModel interface{}, reqbody []byte) (err error) {

	fmt.Printf("%ving: %v\n", method, regUrl)

	req, err := http.NewRequest(method, regUrl.String(), bytes.NewBuffer(reqbody))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.HttpClient.Do(req)
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

func (c *ReflectJsonClient) Get(reqUrl *url.URL, responseModel interface{}) (err error) {
	return c.Do(http.MethodGet, reqUrl, responseModel, nil)
}

func (c *ReflectJsonClient) Post(reqUrl *url.URL, data interface{}, responseModel interface{}) (err error) {
	body, err := json.Marshal(data)
	if err != nil {
		return
	}

	return c.Do(http.MethodPost, reqUrl, responseModel, body)
}
