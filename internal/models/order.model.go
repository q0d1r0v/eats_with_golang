package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID              uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	BranchID        uuid.UUID `gorm:"type:uuid; not null"`
	Branch          Branch    `gorm:"foreignKey:BranchID;references:ID"`
	UserID          uuid.UUID `gorm:"type:uuid;not null;references:ID"`
	User            User      `gorm:"foreignKey:UserID"`
	CourierID       uuid.UUID `gorm:"type:uuid;not null"`
	Courier         Courier   `gorm:"foreignKey:CourierID;references:ID"`
	TotalAmount     float64   `gorm:"not null"`
	DeliveryAddress string    `gorm:"type:text;not null"`
	Status          string    `gorm:"size:50;default:'pending'"`
	gorm.Model
}
