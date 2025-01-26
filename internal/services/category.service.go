package services

import (
	"company_name_eats/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CategoryService provides services related with roles
type CategoryService struct {
	DB *gorm.DB
}

// load all categories
func (s *CategoryService) LoadCategories() ([]models.Category, error) {
	var category []models.Category
	result := s.DB.Find(&category)
	return category, result.Error
}

func (s *CategoryService) CreateCategory(branchID uuid.UUID, name string) (*models.Category, error) {
	category := &models.Category{
		ID:       uuid.New(),
		Name:     name,
		BranchID: branchID,
	}
	if err := s.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
