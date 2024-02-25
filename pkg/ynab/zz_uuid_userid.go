package ynab

import "github.com/google/uuid"

func (id UserID) String() string {
	return (uuid.UUID)(id).String()
}

func (id *UserID) UnmarshalText(b []byte) error {
	return (*uuid.UUID)(id).UnmarshalText(b)
}

func (id UserID) MarshalText() ([]byte, error) {
	return (uuid.UUID)(id).MarshalText()
}

func (id UserID) AsUUID() *uuid.UUID {
	return (*uuid.UUID)(&id)
}

func (id UserID) IsEmpty() bool {
	return (uuid.UUID)(id) == uuid.Nil
}

func ParseUserID(s string) (UserID, error) {
	id, err := uuid.Parse(s)
	return (UserID)(id), err
}

func MustParseUserID(s string) UserID {
	return (UserID)(uuid.MustParse(s))
}
