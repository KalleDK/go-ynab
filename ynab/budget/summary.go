package budget

import "time"

type Summary struct {
	ID           ID         `json:"id"`
	Name         string     `json:"name"`
	LastModified *time.Time `json:"last_modified_on"`
	Settings
}
