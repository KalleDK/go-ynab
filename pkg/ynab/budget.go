package ynab

import "github.com/google/uuid"

// BudgetID identifies the budget from YNAB
type BudgetID = uuid.UUID

// Budget is the model for a budget
type Budget struct {
	ID   BudgetID
	Name string
}
