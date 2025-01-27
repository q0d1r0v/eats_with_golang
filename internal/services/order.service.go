package services

import (
	"company_name_eats/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FoodService provides services related with roles
type OrderService struct {
	DB *gorm.DB
}

// load all orders
func (s *OrderService) LoadAllOrders() ([]models.Order, error) {
	var orders []models.Order
	result := s.DB.Find(&orders)
	return orders, result.Error
}

// create order
func (s *OrderService) CreateOrder(input struct {
	BranchID        uuid.UUID
	UserID          uuid.UUID
	CourierID       uuid.UUID
	DeliveryAddress string
	Items           []struct {
		FoodID    uuid.UUID
		Count     int
		UnitPrice float64
	}
}) (*models.Order, error) {
	order := &models.Order{
		ID:              uuid.New(),
		BranchID:        input.BranchID,
		UserID:          input.UserID,
		CourierID:       input.CourierID,
		DeliveryAddress: input.DeliveryAddress,
		Status:          "pending",
	}

	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := tx.Create(&order).Error; err != nil {
		return nil, err
	}

	var totalAmount float64
	for _, item := range input.Items {
		totalPrice := item.UnitPrice * float64(item.Count)
		orderItem := &models.OrderItem{
			ID:         uuid.New(),
			OrderID:    order.ID,
			FoodID:     item.FoodID,
			Count:      item.Count,
			UnitPrice:  item.UnitPrice,
			TotalPrice: totalPrice,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			return nil, err
		}

		totalAmount += totalPrice
	}

	order.TotalAmount = totalAmount
	if err := tx.Save(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) UpdateOrder(orderID uuid.UUID, status *string) (*models.Order, error) {
	var order models.Order
	if err := s.DB.First(&order, "id = ?", orderID).Error; err != nil {
		return nil, err
	}
	if status != nil {
		order.Status = *status
	}
	if err := s.DB.Save(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
