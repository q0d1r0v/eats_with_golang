package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID       uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Name     string    `gorm:"size:255;not null"`
	BranchID uuid.UUID `gorm:"type:uuid;not null;"`
	Branch   Branch    `gorm:"foreignKey:BranchID"`
	gorm.Model
}
