package client

import (
	"github.com/kalledk/go-ynab/ynab/account"
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

type AccountClient struct {
	Endpoint endpoint.API
}

func (c *AccountClient) Get() (account.Account, error) {
	return account.Get(c.Endpoint)
}

type AccountsClient struct {
	Endpoint endpoint.API
}

func (c *AccountsClient) Get() ([]account.Account, error) {
	return account.GetList(c.Endpoint)
}

func (c *AccountsClient) Account(id account.ID) *AccountClient {
	return &AccountClient{endpoint.Down(c.Endpoint, id.String())}
}
