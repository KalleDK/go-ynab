package ynab

import (
	"encoding/json"
	"fmt"
	"strings"
)

// #region ClearedType

type ClearedType uint8

const (
	NoCleared ClearedType = iota
	Cleared
	Uncleared
	Reconciled
)

var (
	Cleared_jsonname = map[uint8]string{
		1: "cleared",
		2: "uncleared",
		3: "reconciled",
	}
)

var (
	Cleared_name = map[uint8]string{
		0: "NoCleared",
		1: "Cleared",
		2: "Uncleared",
		3: "Reconciled",
	}
)

var (
	Cleared_value = map[string]uint8{
		"cleared": 1,
		"uncleared": 2,
		"reconciled": 3,
	}
)

func (s ClearedType) String() string {
	return Cleared_name[uint8(s)]
}

func (s ClearedType) MarshalJSON() ([]byte, error) {
	if s == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(Cleared_jsonname[uint8(s)])
}

func (s *ClearedType) UnmarshalJSON(data []byte) (err error) {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		*s = ClearedType(0)
	} else if *s, err = ParseClearedType(*value); err != nil {
		return err
	}
	return nil
}

func ParseClearedType(s string) (ClearedType, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return ClearedType(0), nil
	}
	value, ok := Cleared_value[s]
	if !ok {
		return ClearedType(0), fmt.Errorf("%q is not a valid ClearedType", s)
	}
	return ClearedType(value), nil
}

// #endregion
