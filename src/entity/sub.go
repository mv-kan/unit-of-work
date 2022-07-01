package entity

import "github.com/google/uuid"

type Sub struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	Name   string
	TodoID uuid.UUID
}
