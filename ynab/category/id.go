package category

import "github.com/google/uuid"

type ID struct {
	uuid.UUID
}

func (id ID) MarshalString() string {
	var empty uuid.UUID
	if id.UUID == empty {
		return ""
	}

	return id.UUID.String()
}

func NewID(s string) (id ID, err error) {
	id.UUID, err = uuid.Parse(s)
	return
}
