package ynab

import "github.com/google/uuid"

func (id CategoryID) String() string {
	return (uuid.UUID)(id).String()
}

func (id *CategoryID) UnmarshalText(b []byte) error {
	return (*uuid.UUID)(id).UnmarshalText(b)
}

func (id CategoryID) MarshalText() ([]byte, error) {
	return (uuid.UUID)(id).MarshalText()
}

func (id CategoryID) AsUUID() *uuid.UUID {
	return (*uuid.UUID)(&id)
}

func (id CategoryID) IsEmpty() bool {
	return (uuid.UUID)(id) == uuid.Nil
}

func ParseCategoryID(s string) (CategoryID, error) {
	id, err := uuid.Parse(s)
	return (CategoryID)(id), err
}

func MustParseCategoryID(s string) CategoryID {
	return (CategoryID)(uuid.MustParse(s))
}
