package ynab

import (
	"encoding/json"
	"fmt"
	"strings"
)

// #region FlagType

type FlagType uint8

const (
	NoFlag FlagType = iota
	Red
	Orange
	Yellow
	Green
	Blue
	Purple
)

var (
	Flag_jsonname = map[uint8]string{
		1: "red",
		2: "orange",
		3: "yellow",
		4: "green",
		5: "blue",
		6: "purple",
	}
)

var (
	Flag_name = map[uint8]string{
		0: "NoFlag",
		1: "Red",
		2: "Orange",
		3: "Yellow",
		4: "Green",
		5: "Blue",
		6: "Purple",
	}
)

var (
	Flag_value = map[string]uint8{
		"red": 1,
		"orange": 2,
		"yellow": 3,
		"green": 4,
		"blue": 5,
		"purple": 6,
	}
)

func (s FlagType) String() string {
	return Flag_name[uint8(s)]
}

func (s FlagType) MarshalJSON() ([]byte, error) {
	if s == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(Flag_jsonname[uint8(s)])
}

func (s *FlagType) UnmarshalJSON(data []byte) (err error) {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		*s = FlagType(0)
	} else if *s, err = ParseFlagType(*value); err != nil {
		return err
	}
	return nil
}

func ParseFlagType(s string) (FlagType, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return FlagType(0), nil
	}
	value, ok := Flag_value[s]
	if !ok {
		return FlagType(0), fmt.Errorf("%q is not a valid FlagType", s)
	}
	return FlagType(value), nil
}

// #endregion
