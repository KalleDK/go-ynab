package transaction

import "github.com/kalledk/go-ynab/ynab/endpoint"

type DetailWrapper struct {
	Transaction Transaction `json:"transaction"`
}

type SaveTransactionReplyResponse struct {
	Data SaveTransactionReplyWrapper `json:"data"`
}

type SaveTransactionReplyWrapper struct {
	IDs                []ID            `json:"transaction_ids"`
	Transaction        Transaction     `json:"transaction"`
	Transactions       TransactionList `json:"transactions"`
	DuplicateImportIDs []string        `json:"duplicate_import_ids"`
}

type TransactionFormWrapper struct {
	Transaction TransactionForm `json:"transaction,omitempty"`
}

type SaveTransactionListWrapper struct {
	Transactions TransactionFormList `json:"transactions,omitempty"`
}

type DetailResponse struct {
	Data DetailWrapper `json:"data"`
}

func Get(baseEndpoint endpoint.Getter) (t Transaction, err error) {
	var response struct {
		Data struct {
			Transaction `json:"transaction"`
		} `json:"data"`
	}

	err = baseEndpoint.Get(&response)
	if err != nil {
		return
	}

	return response.Data.Transaction, nil
}

type Result struct {
	IDs                []ID        `json:"transaction_ids"`
	Transaction        Transaction `json:"transaction"`
	DuplicateImportIDs []string    `json:"duplicate_import_ids"`
}

type Results struct {
	IDs                []ID        `json:"transaction_ids"`
	Transaction        []Transaction `json:"transactions"`
	DuplicateImportIDs []string    `json:"duplicate_import_ids"`
}

type saveTransaction struct {
	Date   string `json:"date"`
	Amount int64  `json:"amount"`
	Memo   string `json:"memo"`

	PayeeID           payee.ID    `json:"payee_id"`
	PayeeName         string      `json:"payee_name"`
	CategoryID        category.ID `json:"category_id"`
	
	AccountID         account.ID  `json:"account_id"`

	Cleared   ClearingStatus `json:"cleared"`
	FlagColor FlagColor      `json:"flag_color"`
	Approved  bool           `json:"approved"`
	
	ImportID string `json:"import_id"`
}

func makeSaveTransaction(t Transaction) saveTransaction {
	return saveTransaction{
		t.Date,
		t.Amount,
		t.Memo,
		t.PayeeID,
		t.PayeeName,
		t.CategoryID,
		t.AccountID,
		t.Cleared,
		t.FlagColor,
		t.Approved,
		t.ImportID,
	}
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

	err = baseEndpoint.Post(data, &response)
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
		make([]saveTransaction, y),
	}

	for i, t := range ts {
		data.Transactions[i] = makeSaveTransaction(t)
	}

	err = e.Post(data, &response)
	if err != nil {
		return
	}
	return response.Data, nil
}
