package ynab

import "github.com/google/uuid"

type OptUUID interface {
	AsUUID() *uuid.UUID
	IsEmpty() bool
}

func opt_uuid(id OptUUID) *uuid.UUID {
	if id.IsEmpty() {
		return nil
	}
	return id.AsUUID()
}
