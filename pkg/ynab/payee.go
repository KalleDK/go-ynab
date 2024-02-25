//go:generate go run ./ynab_uuid_gen PayeeID

package ynab

import "github.com/google/uuid"

// PayeeID identifies the budget from YNAB
type PayeeID uuid.UUID

// Payee is the model for a payee
type Payee struct {
	ID                PayeeID    `json:"id"`
	Name              string     `json:"name"`
	TransferAccountID *AccountID `json:"transfer_account_id"`
	Deleted           bool       `json:"deleted"`
}

var DefaultPayeeID = PayeeID{}
