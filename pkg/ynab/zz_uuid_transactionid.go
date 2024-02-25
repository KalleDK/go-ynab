package ynab

import "github.com/google/uuid"

func (id TransactionID) String() string {
	return (uuid.UUID)(id).String()
}

func (id *TransactionID) UnmarshalText(b []byte) error {
	return (*uuid.UUID)(id).UnmarshalText(b)
}

func (id TransactionID) MarshalText() ([]byte, error) {
	return (uuid.UUID)(id).MarshalText()
}

func (id TransactionID) AsUUID() *uuid.UUID {
	return (*uuid.UUID)(&id)
}

func (id TransactionID) IsEmpty() bool {
	return (uuid.UUID)(id) == uuid.Nil
}

func ParseTransactionID(s string) (TransactionID, error) {
	id, err := uuid.Parse(s)
	return (TransactionID)(id), err
}

func MustParseTransactionID(s string) TransactionID {
	return (TransactionID)(uuid.MustParse(s))
}
