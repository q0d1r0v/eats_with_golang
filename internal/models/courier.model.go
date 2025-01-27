package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Courier struct {
	ID          uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Fullname    string    `grom:"size:50;not null"`
	Username    string    `gorm:"size:255;not null;unique"`
	Password    string    `gorm:"size:255;not null"`
	PhoneNumber string    `gorm:"size:50;not null"`
	Status      string    `gorm:"size:50;default:'available'"`
	BranchID    uuid.UUID `gorm:"type:uuid;not null;"`
	Branch      Branch    `gorm:"foreignKey:BranchID;references:ID"`
	gorm.Model
}
