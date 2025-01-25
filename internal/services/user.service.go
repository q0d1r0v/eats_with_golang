package services

import (
	"company_name_eats/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// load all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := s.DB.Find(&users)
	return users, result.Error
}
