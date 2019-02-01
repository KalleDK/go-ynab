package transaction

import (
	"encoding/json"
	"log"

	"github.com/kalledk/go-ynab/ynab/endpoint"
)

func Get(baseEndpoint endpoint.Getter) (detail Detail, err error) {
	var response DetailResponse
	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Transaction, nil
}

type SaveData struct {
	Data SaveTransactionWrapper `json:"datas"`
}

func SprintJson(model interface{}) string {
	data, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func Post(baseEndpoint endpoint.API, transaction SaveTransaction) (reply SaveTransactionReplyWrapper, err error) {
	var response SaveTransactionReplyResponse
	data := SaveTransactionWrapper{Transaction: transaction}

	err = baseEndpoint.Post(data, &response)
	if err != nil {
		return
	}
	return response.Data, nil
}
