package services

import (
	"company_name_eats/internal/models"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CourierService provides services related with roles
type CourierService struct {
	DB *gorm.DB
}

// load all couriers
func (s *CourierService) LoadCouriers() ([]models.Courier, error) {
	var couriers []models.Courier
	result := s.DB.Find(&couriers)
	return couriers, result.Error
}

func (s *CourierService) CreateCourier(full_name string, username string, password string, phone_number string, branchID uuid.UUID, status string) (*models.Courier, error) {
	var existingCourier models.Courier
	if err := s.DB.Where("username = ?", username).First(&existingCourier).Error; err == nil {
		return nil, fmt.Errorf("username %s already taken", username)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	courier := &models.Courier{
		ID:          uuid.New(),
		Fullname:    full_name,
		Username:    username,
		Password:    string(hashedPassword),
		PhoneNumber: phone_number,
		BranchID:    branchID,
		Status:      status,
	}
	if err := s.DB.Create(&courier).Error; err != nil {
		return nil, err
	}
	return courier, nil
}

func (s *CourierService) UpdateCourier(courierID uuid.UUID, full_name *string, phone_number *string, branchID *uuid.UUID, status *string) (*models.Courier, error) {
	var courier models.Courier
	if err := s.DB.First(&courier, "id = ?", courierID).Error; err != nil {
		return nil, err
	}
	if full_name != nil {
		courier.Fullname = *full_name
	}
	if full_name != nil {
		courier.Fullname = *full_name
	}
	if phone_number != nil {
		courier.PhoneNumber = *phone_number
	}
	if branchID != nil {
		courier.BranchID = *branchID
	}
	if status != nil {
		courier.Status = *status
	}
	if err := s.DB.Save(&courier).Error; err != nil {
		return nil, err
	}

	return &courier, nil
}

// func (s *RoleService) DeleteRole(roleID uuid.UUID) error {
// 	var role models.Role
// 	if err := s.DB.First(&role, "id = ?", roleID).Error; err != nil {
// 		return err
// 	}
// 	if err := s.DB.Delete(&role).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
