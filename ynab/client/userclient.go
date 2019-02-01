package client

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
	"github.com/kalledk/go-ynab/ynab/user"
)

type UserClient struct {
	Endpoint endpoint.Getter
}

func (c *UserClient) Get() (user.User, error) {
	return user.Get(c.Endpoint)
}
