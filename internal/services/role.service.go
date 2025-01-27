package services

import (
	"company_name_eats/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleService provides services related with roles
type RoleService struct {
	DB *gorm.DB
}

// load all roles
func (s *RoleService) LoadRoles() ([]models.Role, error) {
	var roles []models.Role
	result := s.DB.Find(&roles)
	return roles, result.Error
}

func (s *RoleService) CreateRole(name string) (*models.Role, error) {
	role := &models.Role{
		ID:   uuid.New(),
		Name: name,
	}
	if err := s.DB.Create(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
func (s *RoleService) UpdateRole(roleID uuid.UUID, name *string) (*models.Role, error) {
	var role models.Role
	if err := s.DB.First(&role, "id = ?", roleID).Error; err != nil {
		return nil, err
	}
	if name != nil {
		role.Name = *name
	}
	if err := s.DB.Save(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (s *RoleService) DeleteRole(roleID uuid.UUID) error {
	var role models.Role
	if err := s.DB.First(&role, "id = ?", roleID).Error; err != nil {
		return err
	}
	if err := s.DB.Delete(&role).Error; err != nil {
		return err
	}
	return nil
}
