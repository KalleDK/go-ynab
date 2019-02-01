package api

import "fmt"

type ErrorResponse struct {
	ErrorDetail ErrorDetail `json:"error"`
}

func (er ErrorResponse) Error() string {
	return fmt.Sprintf("%v", er.ErrorDetail.Error())
}
