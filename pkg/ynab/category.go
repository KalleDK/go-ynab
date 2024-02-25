//go:generate go run ./ynab_uuid_gen CategoryID

package ynab

import (
	"encoding/json"

	"github.com/KalleDK/go-money/money"
	"github.com/google/uuid"
)

// CategoryID identifies the budget from YNAB
type CategoryID uuid.UUID

// Category is the model for a category
type Category struct {
	ID                      CategoryID       `json:"id"`
	CategoryGroupID         CategoryGroupID  `json:"category_group_id"`
	CategoryGroupName       string           `json:"category_group_name"`
	Name                    string           `json:"name"`
	Hidden                  bool             `json:"hidden"`
	OriginalCategoryGroupID *CategoryGroupID `json:"original_category_group_id"`
	Note                    *string          `json:"note"`
	Budgeted                Amount           `json:"-"`
	Activity                Amount           `json:"-"`
	Balance                 Amount           `json:"-"`
	GoalType                GoalType         `json:"goal_type"`
	GoalDay                 *int32           `json:"goal_day"`
	GoalCadence             *int32           `json:"goal_cadence"`
	GoalCadenceFrequency    *int32           `json:"goal_cadence_frequency"`
	GoalCreationMonth       *string          `json:"goal_creation_month"`
	GoalTarget              *Amount          `json:"-"`
	GoalTargetMonth         *string          `json:"goal_target_month"`
	Deleted                 bool             `json:"deleted"`
}

func (c *Category) UnmarshalJSON(data []byte) error {
	type Alias Category
	type JSONWrapper struct {
		Alias
		Budgeted   int64  `json:"budgeted"`
		Activity   int64  `json:"activity"`
		Balance    int64  `json:"balance"`
		GoalTarget *int64 `json:"goal_target"`
	}
	wrapper := JSONWrapper{}
	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return err
	}
	*c = Category(wrapper.Alias)
	c.Budgeted = money.FromMilli(wrapper.Budgeted)
	c.Activity = money.FromMilli(wrapper.Activity)
	c.Balance = money.FromMilli(wrapper.Balance)
	if wrapper.GoalTarget != nil {
		value := money.FromMilli(*wrapper.GoalTarget)
		c.GoalTarget = &value
	}
	return nil
}
