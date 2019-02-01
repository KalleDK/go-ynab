package budget

import (
	"time"

	"github.com/kalledk/go-ynab/ynab/budget/settings"
)

type Summary struct {
	ID           ID         `json:"id"`
	Name         string     `json:"name"`
	LastModified *time.Time `json:"last_modified_on"`
	Settings     settings.Settings
}
