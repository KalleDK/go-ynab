package ynab

import "github.com/google/uuid"

// UserID is the ID of a user
type UserID = uuid.UUID

// User is the model of a user
type User struct {
	ID UserID
}
