package client

import (
	"github.com/kalledk/go-ynab/ynab/budget"
	"github.com/kalledk/go-ynab/ynab/budget/settings"
	"github.com/kalledk/go-ynab/ynab/endpoint"
	"github.com/kalledk/go-ynab/ynab/transaction"
)

type BudgetListClient struct {
	Endpoint endpoint.API
}

func (c *BudgetListClient) Budget(id budget.ID) *BudgetClient {
	return &BudgetClient{endpoint.Down(c.Endpoint, id.String())}
}

type BudgetClient struct {
	Endpoint endpoint.API
}

func (c *BudgetClient) Settings() *BudgetSettingsClient {
	return &BudgetSettingsClient{endpoint.Down(c.Endpoint, "settings")}
}

func (c *BudgetClient) Transactions() *TransactionsClient {
	return &TransactionsClient{endpoint.Down(c.Endpoint, "transactions")}
}

func (c *BudgetClient) Payees() *PayeesClient {
	return &PayeesClient{endpoint.Down(c.Endpoint, "payees")}
}

func (c *BudgetClient) Accounts() *AccountsClient {
	return &AccountsClient{endpoint.Down(c.Endpoint, "accounts")}
}

type BudgetSettingsClient struct {
	Endpoint endpoint.API
}

func (c *BudgetSettingsClient) Get() (settings.Settings, error) {
	return settings.Get(c.Endpoint)
}

func GetBudgetSettings(client *Client, id budget.ID) (settings.Settings, error) {
	return client.Budgets().Budget(id).Settings().Get()
}

func GetTransaction(client *BudgetClient, id transaction.ID) (transaction.Transaction, error) {
	return client.Transactions().Transaction(id).Get()
}
