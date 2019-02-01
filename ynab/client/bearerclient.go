package client

import (
	"fmt"
	"net/http"
)

type BearerAuth struct {
	Token string
}

func (a *BearerAuth) AddAuth(header *http.Header) {
	header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}
