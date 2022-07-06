package entity

import "github.com/google/uuid"

type Sub struct {
	Base
	Name   string
	TodoID uuid.UUID
}
