package ynab

import "github.com/google/uuid"

func (id AccountID) String() string {
	return (uuid.UUID)(id).String()
}

func (id *AccountID) UnmarshalText(b []byte) error {
	return (*uuid.UUID)(id).UnmarshalText(b)
}

func (id AccountID) MarshalText() ([]byte, error) {
	return (uuid.UUID)(id).MarshalText()
}

func (id AccountID) AsUUID() *uuid.UUID {
	return (*uuid.UUID)(&id)
}

func (id AccountID) IsEmpty() bool {
	return (uuid.UUID)(id) == uuid.Nil
}

func ParseAccountID(s string) (AccountID, error) {
	id, err := uuid.Parse(s)
	return (AccountID)(id), err
}

func MustParseAccountID(s string) AccountID {
	return (AccountID)(uuid.MustParse(s))
}
