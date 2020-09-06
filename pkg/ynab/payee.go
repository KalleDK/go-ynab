package ynab

import "github.com/google/uuid"

// PayeeID identifies the budget from YNAB
type PayeeID = uuid.UUID

// Payee is the model for a payee
type Payee struct {
	ID   PayeeID
	Name string
}
