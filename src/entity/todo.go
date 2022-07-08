package entity

import "github.com/google/uuid"

type Todo struct {
	Base
	Name   string
	Subs   []Sub
	UserID uuid.UUID
}
