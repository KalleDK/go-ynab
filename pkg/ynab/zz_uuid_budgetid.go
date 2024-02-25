package ynab

import "github.com/google/uuid"

func (id BudgetID) String() string {
	return (uuid.UUID)(id).String()
}

func (id *BudgetID) UnmarshalText(b []byte) error {
	return (*uuid.UUID)(id).UnmarshalText(b)
}

func (id BudgetID) MarshalText() ([]byte, error) {
	return (uuid.UUID)(id).MarshalText()
}

func (id BudgetID) AsUUID() *uuid.UUID {
	return (*uuid.UUID)(&id)
}

func (id BudgetID) IsEmpty() bool {
	return (uuid.UUID)(id) == uuid.Nil
}

func ParseBudgetID(s string) (BudgetID, error) {
	id, err := uuid.Parse(s)
	return (BudgetID)(id), err
}

func MustParseBudgetID(s string) BudgetID {
	return (BudgetID)(uuid.MustParse(s))
}
