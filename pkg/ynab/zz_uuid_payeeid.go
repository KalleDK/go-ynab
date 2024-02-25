package ynab

import "github.com/google/uuid"

func (id PayeeID) String() string {
	return (uuid.UUID)(id).String()
}

func (id *PayeeID) UnmarshalText(b []byte) error {
	return (*uuid.UUID)(id).UnmarshalText(b)
}

func (id PayeeID) MarshalText() ([]byte, error) {
	return (uuid.UUID)(id).MarshalText()
}

func (id PayeeID) AsUUID() *uuid.UUID {
	return (*uuid.UUID)(&id)
}

func (id PayeeID) IsEmpty() bool {
	return (uuid.UUID)(id) == uuid.Nil
}

func ParsePayeeID(s string) (PayeeID, error) {
	id, err := uuid.Parse(s)
	return (PayeeID)(id), err
}

func MustParsePayeeID(s string) PayeeID {
	return (PayeeID)(uuid.MustParse(s))
}
