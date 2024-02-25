//go:generate go run ./ynab_uuid_gen TransactionID

package ynab

import (
	"encoding/json"
	"time"

	"github.com/KalleDK/go-money/money"
	"github.com/google/uuid"
)

// TransactionID identifies the budget from YNAB
type TransactionID uuid.UUID

type NewTransaction struct {
	AccountID       AccountID           `json:"account_id"`                // Required
	Date            time.Time           `json:"date"`                      // Required
	Amount          Amount              `json:"-"`                         // Required; Custom JSON
	PayeeID         PayeeID             `json:"-"`                         // Optional; Custom JSON
	PayeeName       string              `json:"payee_name,omitempty"`      // Optional
	CategoryID      CategoryID          `json:"-"`                         // Optional; Custom JSON
	Memo            string              `json:"memo,omitempty"`            // Optional
	Cleared         ClearedType         `json:"cleared,omitempty"`         // Optional; Custom JSON
	Approved        bool                `json:"approved"`                  // Optional
	FlagColor       FlagType            `json:"flag_color,omitempty"`      // Optional; Custom JSON
	SubTransactions []NewSubTransaction `json:"subtransactions,omitempty"` // Optional
	ImportID        string              `json:"import_id,omitempty"`       // Optional
}

// MarshalJSON converts the struct to json
func (t NewTransaction) MarshalJSON() ([]byte, error) {
	type Alias NewTransaction
	type JSONWrapper struct {
		Alias
		Amount     int64      `json:"amount"`
		PayeeID    *uuid.UUID `json:"payee_id,omitempty"`
		CategoryID *uuid.UUID `json:"category_id,omitempty"`
	}

	wrapper := JSONWrapper{
		Alias:      Alias(t),
		Amount:     t.Amount.AsMilli(),
		PayeeID:    opt_uuid(t.PayeeID),
		CategoryID: opt_uuid(t.CategoryID),
	}

	return json.Marshal(wrapper)

}

type NewSubTransaction struct {
	Amount     Amount     `json:"-"` // Custom unmarshal
	PayeeID    PayeeID    `json:"-"` // Custom unmarshal
	PayeeName  string     `json:"payee_name,omitempty"`
	CategoryID CategoryID `json:"-"` // Custom unmarshal
	Memo       string     `json:"memo,omitempty"`
}

// MarshalJSON converts the struct to json
func (t NewSubTransaction) MarshalJSON() ([]byte, error) {
	type Alias NewSubTransaction
	type JSONWrapper struct {
		Alias
		Amount     int64
		PayeeID    *uuid.UUID `json:"payee_id,omitempty"`
		CategoryID *uuid.UUID `json:"category_id,omitempty"`
	}

	wrapper := JSONWrapper{
		Alias:      Alias(t),
		Amount:     t.Amount.AsMilli(),
		PayeeID:    opt_uuid(t.PayeeID),
		CategoryID: opt_uuid(t.CategoryID),
	}

	return json.Marshal(wrapper)
}

type SubTransaction struct {
	ID                    TransactionID  `json:"id"`
	TransactionID         TransactionID  `json:"transaction_id"`
	Amount                Amount         `json:"-"` // Custom unmarshal
	Memo                  *string        `json:"memo"`
	PayeeID               *PayeeID       `json:"payee_id"`
	PayeeName             *string        `json:"payee_name"`
	CategoryID            *CategoryID    `json:"category_id"`
	CategoryName          *string        `json:"category_name"`
	TransferAccountID     *AccountID     `json:"transfer_account_id"`
	TransferTransactionID *TransactionID `json:"transfer_transaction_id"`
	Deleted               bool           `json:"deleted"`
}

func (t *SubTransaction) UnmarshalJSON(data []byte) error {
	type Alias SubTransaction
	type JSONWrapper struct {
		Alias
		Amount int64
	}

	var wrapper JSONWrapper

	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return err
	}

	wrapper.Alias.Amount = money.FromMilli(wrapper.Amount)

	*t = SubTransaction(wrapper.Alias)

	return nil
}

// Transaction is a transaction on the YNAB site
type Transaction struct {
	ID                      TransactionID    `json:"id"`
	Date                    time.Time        `json:"-"` // Custom unmarshal
	Amount                  Amount           `json:"-"` // Custom unmarshal
	Memo                    *string          `json:"memo"`
	Cleared                 ClearedType      `json:"cleared"`
	Approved                bool             `json:"approved"`
	FlagColor               FlagType         `json:"flag_color"`
	FlagName                *string          `json:"flag_name"`
	AccountID               AccountID        `json:"account_id"`
	PayeeID                 *PayeeID         `json:"payee_id"`
	CategoryID              *CategoryID      `json:"category_id"`
	TransferAccountID       *AccountID       `json:"transfer_account_id"`
	TransferTransactionID   *TransactionID   `json:"transfer_transaction_id"`
	MatchedTransactionID    *TransactionID   `json:"matched_transaction_id"`
	ImportID                *string          `json:"import_id"`
	ImportPayeeName         *string          `json:"import_payee_name"`
	ImportPayeeNameOriginal *string          `json:"import_payee_name_original"`
	DebtTransactionType     *string          `json:"debt_transaction_type"`
	Deleted                 bool             `json:"deleted"`
	AccountName             string           `json:"account_name"`
	PayeeName               *string          `json:"payee_name"`
	CategoryName            *string          `json:"category_name"`
	SubTransactions         []SubTransaction `json:"subtransactions"`
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	type JSONWrapper struct {
		Alias
		Amount int64
		Date   string
	}

	var wrapper JSONWrapper

	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return err
	}

	wrapper.Alias.Amount = money.FromMilli(wrapper.Amount)
	wrapper.Alias.Date, err = time.Parse("2006-01-02", wrapper.Date)
	if err != nil {
		return err
	}

	*t = Transaction(wrapper.Alias)

	return nil
}
