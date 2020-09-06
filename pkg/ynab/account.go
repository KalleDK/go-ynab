package ynab

import "github.com/google/uuid"

// AccountID identifies the budget from YNAB
type AccountID = uuid.UUID

// Account is the model for an Account
type Account struct {
	ID   AccountID
	Name string
}
