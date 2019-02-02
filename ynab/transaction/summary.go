package transaction

import (
	"github.com/kalledk/go-ynab/ynab/account"
	"github.com/kalledk/go-ynab/ynab/category"
	"github.com/kalledk/go-ynab/ynab/payee"
)

type SubTransaction struct {
	Amount            int64      `json:"amount"`
	Memo              string     `json:"memo"`
	PayeeID           payee.ID   `json:"payee_id"`
	CategoryID        string     `json:"category_id"`
	ID                ID         `json:"id"`
	Deleted           bool       `json:"deleted"`
	TransferAccountID account.ID `json:"transfer_account_id"`
	SuperID           ID         `json:"transaction_id"`
}

type Summary struct {
	ID ID `json:"id"`

	Date   string `json:"date"`
	Amount int64  `json:"amount"`

	Memo string `json:"memo"`

	PayeeID    payee.ID    `json:"payee_id"`
	CategoryID category.ID `json:"category_id"`
	AccountID  account.ID  `json:"account_id"`

	Cleared   ClearingStatus `json:"cleared"`
	FlagColor FlagColor      `json:"flag_color"`
	Approved  bool           `json:"approved"`

	ImportID string `json:"import_id"`

	Deleted           bool       `json:"deleted"`
	TransferAccountID account.ID `json:"transfer_account_id"`
}

type Transaction struct {
	ID     ID     `json:"id"`
	Date   string `json:"date"`
	Amount int64  `json:"amount"`
	Memo   string `json:"memo"`

	PayeeID           payee.ID    `json:"payee_id"`
	PayeeName         string      `json:"payee_name"`
	CategoryID        category.ID `json:"category_id"`
	CategoryName      string      `json:"category_name"`
	AccountID         account.ID  `json:"account_id"`
	AccountName       string      `json:"account_name"`
	TransferAccountID account.ID  `json:"transfer_account_id"`

	Cleared   ClearingStatus `json:"cleared"`
	FlagColor FlagColor      `json:"flag_color"`
	Approved  bool           `json:"approved"`
	Deleted   bool           `json:"deleted"`

	ImportID string `json:"import_id"`

	SubTransactions SubTransactionList `json:"subtransactions"`
}

type SubTransactionList = []SubTransaction

type TransactionList = []Transaction
