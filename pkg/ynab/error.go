package ynab

import (
	"fmt"
)

type Error struct {
	ID     string
	Name   string
	Detail string
}

func (e Error) Error() string {
	return fmt.Sprintf("ID: %s, Name: %s, Detail: %s", e.ID, e.Name, e.Detail)
}
