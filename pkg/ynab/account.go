//go:generate go run ./ynab_uuid_gen AccountID

package ynab

import (
	"encoding/json"
	"time"

	"github.com/KalleDK/go-money/money"
	"github.com/google/uuid"
)

// AccountID identifies the account from YNAB
type AccountID uuid.UUID

// TransferPayeeID identifies the account from YNAB
// when used as a transfer payee
type TransferPayeeID = PayeeID

// Account is the model for an Account
type Account struct {
	ID                  AccountID         `json:"id"`
	Name                string            `json:"name"`
	Type                AccountType       `json:"type"`
	OnBudget            bool              `json:"on_budget"`
	Closed              bool              `json:"closed"`
	Note                *string           `json:"note"`
	Balance             Amount            `json:"-"`
	ClearedBalance      Amount            `json:"-"`
	UnclearedBalance    Amount            `json:"-"`
	TransferPayeeID     TransferPayeeID   `json:"transfer_payee_id"`
	DirectImportLinked  bool              `json:"direct_import_linked"`
	DirectImportInError bool              `json:"direct_import_in_error"`
	LastReconsiledAt    *time.Time        `json:"last_reconciled_at"`
	DebtOriginalBalance *Amount           `json:"-"`
	DebtInterestRates   map[string]Amount `json:"-"`
	DebtMinimumPayment  map[string]Amount `json:"-"`
	DebtEscrowAmounts   map[string]Amount `json:"-"`
	Deleted             bool              `json:"deleted"`
}

func (a *Account) UnmarshalJSON(data []byte) error {
	type Alias Account
	type JSONWrapper struct {
		Alias
		Balance             int64            `json:"balance"`
		ClearedBalance      int64            `json:"cleared_balance"`
		UnclearedBalance    int64            `json:"uncleared_balance"`
		DebtOriginalBalance *int64           `json:"debt_original_balance"`
		DebtInterestRates   map[string]int64 `json:"debt_interest_rates"`
		DebtMinimumPayment  map[string]int64 `json:"debt_minimum_payments"`
		DebtEscrowAmounts   map[string]int64 `json:"debt_escrow_amounts"`
	}
	wrapper := JSONWrapper{}
	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return err
	}
	*a = Account(wrapper.Alias)
	a.Balance = money.FromMilli(wrapper.Balance)
	a.ClearedBalance = money.FromMilli(wrapper.ClearedBalance)
	a.UnclearedBalance = money.FromMilli(wrapper.UnclearedBalance)
	if wrapper.DebtOriginalBalance != nil {
		value := money.FromMilli(*wrapper.DebtOriginalBalance)
		a.DebtOriginalBalance = &value
	}
	if wrapper.DebtInterestRates != nil {
		a.DebtInterestRates = make(map[string]Amount, len(wrapper.DebtInterestRates))
		for k, v := range wrapper.DebtInterestRates {
			value := money.FromMilli(v)
			a.DebtInterestRates[k] = value
		}
	}
	if wrapper.DebtMinimumPayment != nil {
		a.DebtMinimumPayment = make(map[string]Amount, len(wrapper.DebtMinimumPayment))
		for k, v := range wrapper.DebtMinimumPayment {
			value := money.FromMilli(v)
			a.DebtMinimumPayment[k] = value
		}
	}
	if wrapper.DebtEscrowAmounts != nil {
		a.DebtEscrowAmounts = make(map[string]Amount, len(wrapper.DebtEscrowAmounts))
		for k, v := range wrapper.DebtEscrowAmounts {
			value := money.FromMilli(v)
			a.DebtEscrowAmounts[k] = value
		}
	}

	return nil
}
