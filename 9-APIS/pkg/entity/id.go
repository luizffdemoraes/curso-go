package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, _ := uuid.Parse(s)
	return ID(id), nil
}
