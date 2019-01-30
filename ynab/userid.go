package ynab

import (
    "github.com/google/uuid"
)

type UserID struct {
	uuid.UUID
}

func NewUserID(s string) (userID UserID, err error) {
    userID.UUID, err = uuid.Parse(s)
    if err != nil {
        return
    }

    return
}
