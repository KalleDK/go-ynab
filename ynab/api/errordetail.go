package api

import "fmt"

type ErrorDetail struct {
	ID     string
	Name   string
	Detail string
}

func (ed ErrorDetail) Error() string {
	return fmt.Sprintf("ID: %v Name: %v Detail: %v", ed.ID, ed.Name, ed.Detail)
}
