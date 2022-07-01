package entity

import "github.com/google/uuid"

type Sub struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name   string
	TodoID uuid.UUID
}
