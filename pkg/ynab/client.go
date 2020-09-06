package ynab

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"path"
)

// Client communicates with the YNAB api
type Client struct {
	scheme  string
	token   Token
	baseURL string
	client  HTTPClient
}

func (c *Client) url(parts ...string) string {
	return c.scheme + path.Join(c.baseURL, path.Join(parts...))

}

func (c *Client) do(method string, url string, payload interface{}, v interface{}) error {

	body, err := func(payload interface{}) (io.Reader, error) {
		if payload == nil {
			return nil, nil
		}

		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		return bytes.NewBuffer(data), nil
	}(payload)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	request.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + string(c.token)},
	}

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		resp := struct {
			Error Error
		}{}

		if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
			return err
		}

		return resp.Error
	}

	resp := struct {
		Data interface{} `json:"data"`
	}{
		Data: v,
	}

	return json.NewDecoder(response.Body).Decode(&resp)

}

// GetUser returns this user
func (c *Client) GetUser() (User, error) {

	path := c.url("user")

	var resp struct {
		User User
	}

	err := c.do(http.MethodGet, path, nil, &resp)
	if err != nil {
		return User{}, err
	}

	return resp.User, nil
}

// GetBudgets returns all Budgets for this user
func (c *Client) GetBudgets() ([]Budget, error) {

	path := c.url("budgets?include_accounts=false")

	var resp struct {
		Budgets []Budget
	}

	err := c.do(http.MethodGet, path, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Budgets, nil
}

// GetPayees returns all Payees for this budget
func (c *Client) GetPayees(budgetID BudgetID) ([]Payee, error) {

	path := c.url("budgets", budgetID.String(), "payees")

	var resp struct {
		Payees []Payee
	}

	err := c.do(http.MethodGet, path, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Payees, nil
}

// GetAccounts return all accounts for the budget
func (c *Client) GetAccounts(budgetID BudgetID) ([]Account, error) {

	path := c.url("budgets", budgetID.String(), "accounts")

	var response struct {
		Accounts []Account
	}

	err := c.do(http.MethodGet, path, nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Accounts, nil
}

// GetCategories return all categories for the budget
func (c *Client) GetCategories(budgetID BudgetID) ([]CategoryGroup, error) {

	path := c.url("budgets", budgetID.String(), "categories")

	var resp struct {
		CategoryGroups []CategoryGroup `json:"category_groups"`
	}

	err := c.do(http.MethodGet, path, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.CategoryGroups, nil
}

// PostTransactions send transactions to YNAB
func (c *Client) PostTransactions(budgetID BudgetID, transactions []Transaction) error {

	path := c.url("budgets", budgetID.String(), "transactions")

	var resp struct {
		TransactionIDs []string `json:"transaction_ids"`
		Dublicates     []string `json:"duplicate_import_ids"`
	}

	req := struct {
		Transaction  *Transaction  `json:"transaction,omitempty"`
		Transactions []Transaction `json:"transactions,omitempty"`
	}{
		Transactions: transactions,
	}

	err := c.do(http.MethodPost, path, req, &resp)
	if err != nil {
		return err
	}

	return nil
}
