package transaction

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kalledk/go-ynab/ynab/account"
	"github.com/kalledk/go-ynab/ynab/endpoint"
)

type Result struct {
	IDs                []ID        `json:"transaction_ids"`
	Transaction        Transaction `json:"transaction"`
	DuplicateImportIDs []string    `json:"duplicate_import_ids"`
}

type Results struct {
	IDs                []ID          `json:"transaction_ids"`
	Transactions       []Transaction `json:"transactions"`
	DuplicateImportIDs []string      `json:"duplicate_import_ids"`
}

type saveTransaction struct {
	Date   string `json:"date"`
	Amount int64  `json:"amount"`
	Memo   string `json:"memo,omitempty"`

	PayeeID    string `json:"payee_id,omitempty"`
	PayeeName  string `json:"payee_name,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	AccountID account.ID `json:"account_id"`

	Cleared   ClearingStatus `json:"cleared,omitempty"`
	FlagColor FlagColor      `json:"flag_color,omitempty"`
	Approved  bool           `json:"approved,omitempty"`

	ImportID string `json:"import_id,omitempty"`
}

func makeSaveTransaction(t Transaction) saveTransaction {

	return saveTransaction{
		t.Date,
		t.Amount,
		t.Memo,
		t.PayeeID.MarshalString(),
		t.PayeeName,
		t.CategoryID.MarshalString(),
		t.AccountID,
		t.Cleared,
		t.FlagColor,
		t.Approved,
		t.ImportID,
	}
}

func Get(baseEndpoint endpoint.Getter) (t Transaction, err error) {
	var response struct {
		Data struct {
			Transaction Transaction `json:"transaction"`
		} `json:"data"`
	}

	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Transaction, nil
}

func Post(e endpoint.API, t Transaction) (r Result, err error) {
	var response struct {
		Data Result `json:"data"`
	}

	data := struct {
		Transaction saveTransaction `json:"transaction"`
	}{
		makeSaveTransaction(t),
	}

	fmt.Println(SprintJSON(data))

	err = e.Post(data, &response)
	if err != nil {
		return
	}
	return response.Data, nil
}

func PostList(e endpoint.API, ts []Transaction) (reply Results, err error) {
	var response struct {
		Data Results `json:"data"`
	}

	data := struct {
		Transactions []saveTransaction `json:"transactions"`
	}{
		make([]saveTransaction, len(ts)),
	}

	for i, t := range ts {
		data.Transactions[i] = makeSaveTransaction(t)
	}

	fmt.Println(SprintJSON(data))

	err = e.Post(data, &response)
	if err != nil {
		return
	}
	return response.Data, nil
}

func SprintJSON(model interface{}) string {
	data, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
