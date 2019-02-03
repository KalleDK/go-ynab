package user

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

func Get(baseEndpoint endpoint.Getter) (user User, err error) {

	var response struct {
		Data struct {
			User User `json:"user"`
		} `json:"data"`
	}

	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.User, nil
}
