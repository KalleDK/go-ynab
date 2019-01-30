package account

import "github.com/google/uuid"

type ID struct {
    uuid.UUID
}

func NewID(s string) (id ID, err error) {
    id.UUID, err = uuid.Parse(s)
    return
}
