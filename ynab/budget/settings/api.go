package settings

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

type Response struct {
	Data Wrapper `json:"data"`
}

type Wrapper struct {
	Settings Settings `json:"settings"`
}

func Get(endpoint endpoint.Getter) (settings Settings, err error) {
	var response Response

	err = endpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Settings, nil
}
