package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Email    string     `gorm:"size:255;not null;unique"`
	Password string     `gorm:"size:255;not null"`
	BranchID *uuid.UUID `gorm:"type:uuid"`
	Branch   Branch     `gorm:"foreignKey:BranchID;references:ID"`
	RoleID   *uuid.UUID `gorm:"type:uuid"`
	Role     Role       `gorm:"foreignKey:RoleID;references:ID"`
	gorm.Model
}
