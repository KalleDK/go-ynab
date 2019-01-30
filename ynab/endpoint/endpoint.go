//go:generate mockgen -destination mock_endpoint/mock_endpoint.go -source endpoint.go Getter,Poster,Putter

package endpoint

type Getter interface {
	Get(path string, responseModel interface{}) (err error)
}

type Poster interface {
	Post(path string, responseModel interface{}) (err error)
}

type Putter interface {
	Put(path string, responseModel interface{}) (err error)
}
