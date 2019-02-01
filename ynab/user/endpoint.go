package user

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

func Get(baseEndpoint endpoint.Getter) (user User, err error) {
	var response Response
	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.User, nil
}
