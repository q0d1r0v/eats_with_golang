package controllers

import (
	"company_name_eats/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	OrderService *services.OrderService
}
type CreateOrderInput struct {
	BranchID        uuid.UUID
	UserID          uuid.UUID
	CourierID       uuid.UUID
	DeliveryAddress string
	Items           []struct {
		FoodID    uuid.UUID
		Count     int
		UnitPrice float64
	}
}

func (oc *OrderController) LoadAllOrders(c *gin.Context) {
	orders, err := oc.OrderService.LoadAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var orderInput CreateOrderInput
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := oc.OrderService.CreateOrder(struct {
		BranchID        uuid.UUID
		UserID          uuid.UUID
		CourierID       uuid.UUID
		DeliveryAddress string
		Items           []struct {
			FoodID    uuid.UUID
			Count     int
			UnitPrice float64
		}
	}{
		BranchID:        orderInput.BranchID,
		UserID:          orderInput.UserID,
		CourierID:       orderInput.CourierID,
		DeliveryAddress: orderInput.DeliveryAddress,
		Items:           orderInput.Items,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order!"})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (oc *OrderController) UpdateOrder(c *gin.Context) {
	var orderInput struct {
		OrderID uuid.UUID `json:"order_id" binding:"required"`
		Status  string    `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order_id and status are required"})
		return
	}
	order, err := oc.OrderService.UpdateOrder(orderInput.OrderID, &orderInput.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order!"})
		return
	}
	c.JSON(http.StatusOK, order)
}
