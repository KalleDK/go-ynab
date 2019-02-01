package payee

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

func Get(baseEndpoint endpoint.Getter) (payee Payee, err error) {
	var response Response
	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Payee, nil
}

func GetList(baseEndpoint endpoint.Getter) (payees Payees, err error) {
	var response ResponseList
	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Payees, nil
}
