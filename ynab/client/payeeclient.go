package client

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
	"github.com/kalledk/go-ynab/ynab/payee"
)

type PayeeClient struct {
	Endpoint endpoint.API
}

func (c *PayeeClient) Get() (payee.Payee, error) {
	return payee.Get(c.Endpoint)
}

type PayeesClient struct {
	Endpoint endpoint.API
}

func (c *PayeesClient) Get() (payee.Payees, error) {
	return payee.GetList(c.Endpoint)
}

func (c *PayeesClient) Payee(id payee.ID) *PayeeClient {
	return &PayeeClient{endpoint.Down(c.Endpoint, id.String())}
}
