package ynab

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// TransactionID identifies the budget from YNAB
type TransactionID = uuid.UUID

// Transaction is a transaction on the YNAB site
type Transaction struct {
	Account  Account
	Date     time.Time
	Amount   Amount
	Payee    Payee
	Category Category
	Memo     string
	Cleared  string
	Approved bool
	ID       TransactionID
}

// MarshalJSON converts the struct to json
func (t Transaction) MarshalJSON() ([]byte, error) {
	yt := struct {
		AccountID  AccountID  `json:"account_id"`
		Date       time.Time  `json:"date"`
		Amount     int64      `json:"amount"`
		PayeeID    *uuid.UUID `json:"payee_id,omitempty"`
		PayeeName  *string    `json:"payee_name,omitempty"`
		CategoryID uuid.UUID  `json:"category_id"`
		Memo       *string    `json:"memo,omitempty"`
		Cleared    string     `json:"cleared"`
		Approved   bool       `json:"approved"`
		ID         uuid.UUID  `json:"import_id"`
	}{
		AccountID:  t.Account.ID,
		Date:       t.Date,
		Amount:     t.Amount.AsMilli(),
		CategoryID: t.Category.ID,
		Cleared:    t.Cleared,
		Approved:   t.Approved,
		ID:         t.ID,
	}

	if (t.Payee.ID == uuid.UUID{}) {
		yt.PayeeName = &t.Payee.Name
	} else {
		yt.PayeeID = &t.Payee.ID
	}

	if t.Memo != "" {
		yt.Memo = &t.Memo
	}

	return json.Marshal(yt)
}
