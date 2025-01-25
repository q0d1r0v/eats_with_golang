package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID         uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	FoodID     uuid.UUID `gorm:"type:uuid;not null;"`
	Food       Food      `gorm:"foreignKey:FoodID; references:ID"`
	OrderID    uuid.UUID `gorm:"type:uuid;not null;"`
	Order      Order     `gorm:"foreignKey:OrderID; references:ID"`
	Count      int       `gorm:"not null"`
	UnitPrice  float64   `gorm:"not null"`
	TotalPrice float64   `gorm:"not null"`
	gorm.Model
}
