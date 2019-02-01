package budget

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

func Get(baseEndpoint endpoint.Getter) (detail Detail, err error) {
	var response DetailResponse
	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Budget, nil
}

func GetList(baseEndpoint endpoint.Getter) (budgets []Summary, err error) {
	var response SummaryResponse
	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Budgets, nil
}
