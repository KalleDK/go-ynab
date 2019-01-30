package budget

import (
	"path"
)

type APIGetter interface {
	Get(path string, responseModel interface{}) (err error)
}

type Client struct {
	ID        ID
	Endpoint  string
	APIClient APIGetter
}

func NewClient(budgetID ID, apiClient APIGetter) (client *Client) {
	budgetPath := path.Join("budgets", budgetID.String())
	return &Client{
		ID:        budgetID,
		Endpoint:  budgetPath,
		APIClient: apiClient,
	}
}

func (c *Client) Get(subpath string, responseModel interface{}) (err error) {
	endpoint_path := path.Join(c.Endpoint, subpath)
	return c.APIClient.Get(endpoint_path, responseModel)
}

/*
func (c *Client) GetAccount(id account.ID) (a account.Account, err error) {
    return account.Get(c, id)
}
*/

func (c *Client) GetSettings() (settings Settings, err error) {
	var resp SettingsResponse
	err = c.Get("settings", &resp)
	if err != nil {
		return
	}

	settings = resp.Data.Settings
	return
}
