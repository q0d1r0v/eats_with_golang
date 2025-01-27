package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Branch struct {
	ID   uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Name string    `gorm:"size:255;not null"`
	gorm.Model
}
