package ynab

import "github.com/google/uuid"

type CategoryID struct {
	uuid.UUID
}

func (a CategoryID) String() (string) {
	return a.UUID.String()
}

func NewCategoryID(s string) (categoryID CategoryID, err error) {
    u, err := uuid.Parse(s)
    if err != nil {
        return
    }

    return CategoryID{u}, nil
}
