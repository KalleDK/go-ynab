package budget

import "github.com/kalledk/go-ynab/ynab/endpoint"

func Get(baseEndpoint endpoint.Getter) (budget Detail, err error) {
	var response struct {
		Data struct {
			Budget Detail `json:"budget"`
		} `json:"data"`
	}

	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Budget, nil
}

func GetList(baseEndpoint endpoint.Getter) (budgets []Summary, err error) {
	var response struct {
		Data struct {
			Budgets []Summary `json:"budgets"`
		} `json:"data"`
	}

	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Budgets, nil
}
