package ynab

import (
	"encoding/json"
	"fmt"
	"strings"
)

// #region GoalType

type GoalType uint8

const (
	NoGoal GoalType = iota
	TB
	TBD
	MF
	NEED
	DEBT
)

var (
	Goal_jsonname = map[uint8]string{
		1: "TB",
		2: "TBD",
		3: "MF",
		4: "NEED",
		5: "DEBT",
	}
)

var (
	Goal_name = map[uint8]string{
		0: "NoGoal",
		1: "TB",
		2: "TBD",
		3: "MF",
		4: "NEED",
		5: "DEBT",
	}
)

var (
	Goal_value = map[string]uint8{
		"TB": 1,
		"TBD": 2,
		"MF": 3,
		"NEED": 4,
		"DEBT": 5,
	}
)

func (s GoalType) String() string {
	return Goal_name[uint8(s)]
}

func (s GoalType) MarshalJSON() ([]byte, error) {
	if s == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(Goal_jsonname[uint8(s)])
}

func (s *GoalType) UnmarshalJSON(data []byte) (err error) {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		*s = GoalType(0)
	} else if *s, err = ParseGoalType(*value); err != nil {
		return err
	}
	return nil
}

func ParseGoalType(s string) (GoalType, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return GoalType(0), nil
	}
	value, ok := Goal_value[s]
	if !ok {
		return GoalType(0), fmt.Errorf("%q is not a valid GoalType", s)
	}
	return GoalType(value), nil
}

// #endregion
