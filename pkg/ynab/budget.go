//go:generate go run ./ynab_uuid_gen BudgetID

package ynab

import (
	"time"

	"github.com/google/uuid"
)

// BudgetID identifies the budget from YNAB
type BudgetID uuid.UUID

type DateFormat struct {
	Format string `json:"format"`
}

type CurrencyFormat struct {
	IsoCode          string `json:"iso_code"`
	ExampleFormat    string `json:"example_format"`
	DecimalDigits    int    `json:"decimal_digits"`
	DecimalSeparator string `json:"decimal_separator"`
	SymbolFirst      bool   `json:"symbol_first"`
	GroupSeparator   string `json:"group_separator"`
	CurrencySymbol   string `json:"currency_symbol"`
	DisplaySymbol    bool   `json:"display_symbol"`
}

// Budget is the model for a budget
type Budget struct {
	ID             BudgetID       `json:"id"`
	Name           string         `json:"name"`
	LastModifiedOn time.Time      `json:"last_modified_on"`
	FirstMonth     string         `json:"first_month"`
	LastMonth      string         `json:"last_month"`
	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`
}

type BudgetDetail struct {
	Budget
	Accounts       []Account       `json:"accounts"`
	Payees         []Payee         `json:"payees"`
	CategoryGroups []CategoryGroup `json:"category_groups"`
	Categories     []Category      `json:"categories"`
}
