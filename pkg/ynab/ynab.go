package ynab

import (
	"github.com/KalleDK/go-money/money"
)

// Amount is the common type to contain values
type Amount = money.Amount

// Token is the auth token from YNAB
type Token string
