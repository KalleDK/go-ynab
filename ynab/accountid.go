package ynab

import "github.com/google/uuid"

type AccountID struct {
	uuid.UUID
}

func (a AccountID) String() string {
	return a.UUID.String()
}

func NewAccountID(s string) (accountID AccountID, err error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return
	}

	return AccountID{u}, nil
}
