package ynab

import "github.com/google/uuid"

func (id CategoryGroupID) String() string {
	return (uuid.UUID)(id).String()
}

func (id *CategoryGroupID) UnmarshalText(b []byte) error {
	return (*uuid.UUID)(id).UnmarshalText(b)
}

func (id CategoryGroupID) MarshalText() ([]byte, error) {
	return (uuid.UUID)(id).MarshalText()
}

func (id CategoryGroupID) AsUUID() *uuid.UUID {
	return (*uuid.UUID)(&id)
}

func (id CategoryGroupID) IsEmpty() bool {
	return (uuid.UUID)(id) == uuid.Nil
}

func ParseCategoryGroupID(s string) (CategoryGroupID, error) {
	id, err := uuid.Parse(s)
	return (CategoryGroupID)(id), err
}

func MustParseCategoryGroupID(s string) CategoryGroupID {
	return (CategoryGroupID)(uuid.MustParse(s))
}
