package account

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

func Get(baseEndpoint endpoint.Getter) (account Account, err error) {
	var response Response
	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Account, nil
}

func GetList(baseEndpoint endpoint.Getter) (accounts AccountList, err error) {
	var response ResponseList
	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Accounts, nil
}
