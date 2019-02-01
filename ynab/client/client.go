package client

import (
	"net/http"
	"net/url"

	"golang.org/x/oauth2"

	"github.com/kalledk/go-ynab/ynab/api"
	"github.com/kalledk/go-ynab/ynab/endpoint"
	"github.com/kalledk/go-ynab/ynab/user"
)

type Client struct {
	Endpoint endpoint.API
}

var defaultURL, _ = url.Parse("https://api.youneedabudget.com/v1/")

func NewClient(token api.AccessToken) *Client {

	return &Client{
		Endpoint: &APIEndpoint{
			Url: defaultURL,
			JsonClient: &ReflectJsonClient{
				ErrorModel: api.ErrorResponse{},
				HttpClient: &http.Client{
					Transport: &oauth2.Transport{
						Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token.String()}),
					},
				},
			},
		},
	}
}

func (c *Client) User() *UserClient {
	return &UserClient{endpoint.Down(c.Endpoint, "user")}
}

func (c *Client) Budgets() *BudgetListClient {
	return &BudgetListClient{endpoint.Down(c.Endpoint, "budgets")}
}

func GetUser(client *Client) (user.User, error) {
	return client.User().Get()
}
