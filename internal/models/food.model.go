package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Food struct {
	ID         uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Name       string    `gorm:"size:255;not null"`
	Price      float64   `gorm:"not null"`
	CategoryID uuid.UUID `gorm:"type:uuid;not null"`
	Category   Category  `gorm:"foreignKey:CategoryID;references:ID"`
	gorm.Model
}
