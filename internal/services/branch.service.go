package services

import (
	"company_name_eats/internal/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BranchService provides services related with roles
type BranchService struct {
	DB *gorm.DB
}

// load all branches
func (s *BranchService) LoadBranches() ([]models.Branch, error) {
	var branch []models.Branch
	result := s.DB.Find(&branch)
	return branch, result.Error
}

func (s *BranchService) CreateBranch(name string) (*models.Branch, error) {
	branch := &models.Branch{
		ID:   uuid.New(),
		Name: name,
	}

	if err := s.DB.Create(&branch).Error; err != nil {
		return nil, err
	}

	return branch, nil
}

func (s *BranchService) UpdateBranchNameByID(branchID uuid.UUID, newName string) (*models.Branch, error) {
	branch := &models.Branch{}
	err := s.DB.First(&branch, "id = ?", branchID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("branch not found")
	} else if err != nil {
		return nil, err
	}
	branch.Name = newName
	err = s.DB.Save(&branch).Error
	if err != nil {
		return nil, err
	}
	return branch, nil
}
func (s *BranchService) DeleteBranchByID(branchID uuid.UUID) error {
	branch := &models.Branch{}
	err := s.DB.First(&branch, "id = ?", branchID).Error
	if err == gorm.ErrRecordNotFound {
		return fmt.Errorf("branch not found")
	} else if err != nil {
		return err
	}
	err = s.DB.Delete(&branch).Error
	if err != nil {
		return err
	}

	return nil
}
