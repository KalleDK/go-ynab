package ynab

import (
	"encoding/json"
	"fmt"
	"strings"
)

// #region AccountType

type AccountType uint8

const (
	NoAccount AccountType = iota
	Checking
	Savings
	Cash
	CreditCard
	LineOfCredit
	OtherAsset
	OtherLiability
	Mortgage
	AutoLoan
	StudentLoan
	PersonalLoan
	MedicalDebt
	OtherDebt
)

var (
	Account_jsonname = map[uint8]string{
		1: "checking",
		2: "savings",
		3: "cash",
		4: "creditCard",
		5: "lineOfCredit",
		6: "otherAsset",
		7: "otherLiability",
		8: "mortgage",
		9: "autoLoan",
		10: "studentLoan",
		11: "personalLoan",
		12: "medicalDebt",
		13: "otherDebt",
	}
)

var (
	Account_name = map[uint8]string{
		0: "NoAccount",
		1: "Checking",
		2: "Savings",
		3: "Cash",
		4: "CreditCard",
		5: "LineOfCredit",
		6: "OtherAsset",
		7: "OtherLiability",
		8: "Mortgage",
		9: "AutoLoan",
		10: "StudentLoan",
		11: "PersonalLoan",
		12: "MedicalDebt",
		13: "OtherDebt",
	}
)

var (
	Account_value = map[string]uint8{
		"checking": 1,
		"savings": 2,
		"cash": 3,
		"creditCard": 4,
		"lineOfCredit": 5,
		"otherAsset": 6,
		"otherLiability": 7,
		"mortgage": 8,
		"autoLoan": 9,
		"studentLoan": 10,
		"personalLoan": 11,
		"medicalDebt": 12,
		"otherDebt": 13,
	}
)

func (s AccountType) String() string {
	return Account_name[uint8(s)]
}

func (s AccountType) MarshalJSON() ([]byte, error) {
	if s == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(Account_jsonname[uint8(s)])
}

func (s *AccountType) UnmarshalJSON(data []byte) (err error) {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		*s = AccountType(0)
	} else if *s, err = ParseAccountType(*value); err != nil {
		return err
	}
	return nil
}

func ParseAccountType(s string) (AccountType, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return AccountType(0), nil
	}
	value, ok := Account_value[s]
	if !ok {
		return AccountType(0), fmt.Errorf("%q is not a valid AccountType", s)
	}
	return AccountType(value), nil
}

// #endregion
