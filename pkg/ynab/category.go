package ynab

import "github.com/google/uuid"

// CategoryID identifies the budget from YNAB
type CategoryID = uuid.UUID

// Category is the model for a category
type Category struct {
	ID     CategoryID
	Name   string
	Hidden bool
}

// CategoryGroup is a collection of Categories
type CategoryGroup struct {
	ID         CategoryID
	Name       string
	Categories []Category
}
