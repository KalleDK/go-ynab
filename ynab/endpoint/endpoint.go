//go:generate mockgen -destination mock_endpoint/mock_endpoint.go -source endpoint.go Getter,Poster,API

package endpoint

import (
	"path"
)

type Endpoint struct {
	Base API
	Path string
}

func (e *Endpoint) GetVia(subpath string, responseModel interface{}) (err error) {
	return e.Base.GetVia(path.Join(e.Path, subpath), responseModel)
}

func (e *Endpoint) Get(responseModel interface{}) (err error) {
	return e.GetVia("/", responseModel)
}

func (e *Endpoint) PostVia(subpath string, data interface{}, responseModel interface{}) (err error) {
	return e.Base.PostVia(path.Join(e.Path, subpath), data, responseModel)
}

func (e *Endpoint) Post(data interface{}, responseModel interface{}) (err error) {
	return e.PostVia("/", data, responseModel)
}

func Down(base API, path string) *Endpoint {
	return &Endpoint{base, path}
}

type API interface {
	Getter
	Poster
}

type Getter interface {
	Get(responseModel interface{}) (err error)
	GetVia(path string, responseModel interface{}) (err error)
}

type Poster interface {
	Post(data interface{}, responseModel interface{}) (err error)
	PostVia(path string, data interface{}, responseModel interface{}) (err error)
}
