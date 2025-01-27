package services

import (
	"company_name_eats/internal/models"

	"github.com/google/uuid"
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

func (s *FoodService) CreateFood(CategoryID uuid.UUID, name string, price float64) (*models.Food, error) {
	food := &models.Food{
		ID:         uuid.New(),
		CategoryID: CategoryID,
		Name:       name,
		Price:      price,
	}
	if err := s.DB.Create(&food).Error; err != nil {
		return nil, err
	}
	return food, nil
}

func (s *FoodService) UpdateFood(foodID uuid.UUID, categoryID *uuid.UUID, name *string, price *float64) (*models.Food, error) {
	var food models.Food
	if err := s.DB.First(&food, "id = ?", foodID).Error; err != nil {
		return nil, err
	}
	if categoryID != nil {
		food.CategoryID = *categoryID
	}
	if name != nil {
		food.Name = *name
	}
	if price != nil {
		food.Price = *price
	}
	if err := s.DB.Save(&food).Error; err != nil {
		return nil, err
	}

	return &food, nil
}

func (s *FoodService) DeleteFood(foodID uuid.UUID) error {
	var food models.Food
	if err := s.DB.First(&food, "id = ?", foodID).Error; err != nil {
		return err
	}
	if err := s.DB.Delete(&food).Error; err != nil {
		return err
	}
	return nil
}
