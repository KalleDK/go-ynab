package transaction

import (
	"github.com/kalledk/go-ynab/ynab/account"
	"github.com/kalledk/go-ynab/ynab/payee"
)

type BaseTransaction struct {
	ID ID `json:"id"`

	// Amount sub-transaction amount in milliunits format
	Amount int64 `json:"amount"`
	// Deleted Deleted sub-transactions will only be included in delta requests.
	Deleted bool `json:"deleted"`

	Memo       string   `json:"memo"`
	PayeeID    payee.ID `json:"payee_id"`
	CategoryID string   `json:"category_id"`
	// TransferAccountID If a transfer, the account_id which the
	// sub-transaction transfers to
	TransferAccountID string `json:"transfer_account_id"`
}

type SubTransaction struct {
	BaseTransaction
	TransactionID string `json:"transaction_id"`
}

type Summary struct {
	BaseTransaction
	Date string `json:"date"`

	// Amount Transaction amount in milliunits format

	Cleared   ClearingStatus `json:"cleared"`
	Approved  bool           `json:"approved"`
	AccountID account.ID     `json:"account_id"`

	FlagColor FlagColor `json:"flag_color"`

	// ImportID If the Transaction was imported, this field is a unique (by account) import
	// identifier. If this transaction was imported through File Based Import or
	// Direct Import and not through the API, the import_id will have the format:
	// 'YNAB:[milliunit_amount]:[iso_date]:[occurrence]'. For example, a transaction
	// dated 2015-12-30 in the amount of -$294.23 USD would have an import_id of
	// 'YNAB:-294230:2015-12-30:1’. If a second transaction on the same account
	// was imported and had the same date and same amount, its import_id would
	// be 'YNAB:-294230:2015-12-30:2’.
	ImportID string `json:"import_id"`
}
