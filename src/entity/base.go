package entity

import (
	"github.com/google/uuid"
)

// Base contains common columns for all tables.
type Base struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
}

// // BeforeCreate will set a UUID rather than numeric ID.
// func (base *Base) BeforeCreate(tx *gorm.DB) error {
// 	empty := uuid.UUID{}
// 	if base.ID == empty {
// 		id, err := uuid.NewUUID()
// 		if err != nil {
// 			return err
// 		}
// 		base.ID = id
// 	}
// 	return nil
// }
