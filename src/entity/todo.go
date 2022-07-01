package entity

import "github.com/google/uuid"

type Todo struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
	Subs []Sub
}
