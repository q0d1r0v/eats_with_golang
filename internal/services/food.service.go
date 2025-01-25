package services

import (
	"company_name_eats/internal/models"

	"gorm.io/gorm"
)

// FoodService provides services related with roles
type FoodService struct {
	DB *gorm.DB
}

// load all foods
func (s *FoodService) LoadFoods() ([]models.Food, error) {
	var foods []models.Food
	result := s.DB.Find(&foods)
	return foods, result.Error
}
