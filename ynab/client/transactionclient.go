package client

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
	"github.com/kalledk/go-ynab/ynab/transaction"
)

type TransactionsClient struct {
	Endpoint endpoint.API
}

func (c *TransactionsClient) Add(t transaction.SaveTransaction) (transaction.SaveTransactionReplyWrapper, error) {
	return transaction.Post(c.Endpoint, t)
}

func (c *TransactionsClient) Transaction(id transaction.ID) *TransactionClient {
	return &TransactionClient{endpoint.Down(c.Endpoint, id.String())}
}

type TransactionClient struct {
	Endpoint endpoint.API
}

func (c *TransactionClient) Get() (transaction.Detail, error) {
	return transaction.Get(c.Endpoint)
}
