package services

import (
	"company_name_eats/internal/models"

	"gorm.io/gorm"
)

// RoleService provides services related with roles
type RoleService struct {
	DB *gorm.DB
}

// load all roles
func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	result := s.DB.Find(&roles)
	return roles, result.Error
}
