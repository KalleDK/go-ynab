//go:generate go run ./ynab_uuid_gen CategoryGroupID

package ynab

import (
	"github.com/google/uuid"
)

type CategoryGroupID uuid.UUID

// CategoryGroup is a collection of Categories
type CategoryGroup struct {
	ID         CategoryGroupID `json:"id"`
	Name       string          `json:"name"`
	Hidden     bool            `json:"hidden"`
	Deleted    bool            `json:"deleted"`
	Categories []Category      `json:"categories"`
}
