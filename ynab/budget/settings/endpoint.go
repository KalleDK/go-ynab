package settings

import (
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

func Get(endpoint endpoint.Getter) (settings Settings, err error) {
	var response Response

	err = endpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Settings, nil
}
