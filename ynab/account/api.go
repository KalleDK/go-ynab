package account

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

func Get(e endpoint.Getter) (account Account, err error) {
	var response struct {
		Data struct {
			Account Account `json:"account"`
		} `json:"data"`
	}

	err = e.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Account, nil
}

func GetList(e endpoint.Getter) (accounts []Account, err error) {
	var response struct {
		Data struct {
			Accounts []Account `json:"accounts"`
		} `json:"data"`
	}

	err = e.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Accounts, nil
}
