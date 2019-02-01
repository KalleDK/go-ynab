package settings

import "github.com/kalledk/go-ynab/ynab"

type Settings struct {
	DateFormat     ynab.DateFormat     `json:"date_format"`
	CurrencyFormat ynab.CurrencyFormat `json:"currency_format"`
}
