package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Courier struct {
	ID          uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Fullname    string    `grom:"size:50;not null"`
	PhoneNumber string    `gorm:"size:50;not null"`
	Status      string    `gorm:"size:50;default:'available'"`
	gorm.Model
}
