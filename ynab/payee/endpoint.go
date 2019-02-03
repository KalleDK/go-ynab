package payee

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

func Get(baseEndpoint endpoint.Getter) (payee Payee, err error) {
	var response struct {
		Data struct {
			Payee Payee `json:"payee"`
		} `json:"data"`
	}

	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Payee, nil
}

func GetList(baseEndpoint endpoint.Getter) (payees []Payee, err error) {
	var response struct {
		Data struct {
			Payees []Payee `json:"payees"`
		} `json:"data"`
	}

	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Payees, nil
}
