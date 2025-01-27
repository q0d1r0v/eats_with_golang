package controllers

import (
	"company_name_eats/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CourierController struct {
	CourierService *services.CourierService
}

func (cc *CourierController) LoadCouriers(c *gin.Context) {
	couriers, err := cc.CourierService.LoadCouriers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch couriers"})
		return
	}
	c.JSON(http.StatusOK, couriers)
}

func (cc *CourierController) CreateCourier(c *gin.Context) {
	var courierInput struct {
		FullName    string    `json:"full_name" binding:"required"`
		Username    string    `json:"username" binding:"required"`
		Password    string    `json:"password" binding:"required"`
		PhoneNumber string    `json:"phone_number" binding:"required"`
		BranchID    uuid.UUID `json:"branch_id" binding:"required"`
		Status      string    `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&courierInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "full_name, username, password phone_number, branch_id and status are required"})
		return
	}
	courier, err := cc.CourierService.CreateCourier(courierInput.FullName, courierInput.Username, courierInput.Password, courierInput.PhoneNumber, courierInput.BranchID, courierInput.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create courier!"})
		return
	}
	c.JSON(http.StatusCreated, courier)
}

func (cc *CourierController) UpdateCourier(c *gin.Context) {
	var courierInput struct {
		CourierID   uuid.UUID `json:"courier_id" binding:"required"`
		FullName    string    `json:"full_name" binding:"required"`
		PhoneNumber string    `json:"phone_number" binding:"required"`
		BranchID    uuid.UUID `json:"branch_id" binding:"required"`
		Status      string    `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&courierInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "courier_id, full_name, phone_number, branch_id and status are required"})
		return
	}
	courier, err := cc.CourierService.UpdateCourier(courierInput.CourierID, &courierInput.FullName, &courierInput.PhoneNumber, &courierInput.BranchID, &courierInput.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update courier!"})
		return
	}
	c.JSON(http.StatusOK, courier)
}

// func (rc *RoleController) DeleteRole(c *gin.Context) {
// 	var roleInput struct {
// 		RoleID uuid.UUID `json:"role_id" binding:"required"`
// 	}
// 	if err := c.ShouldBindJSON(&roleInput); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "role_id is required"})
// 		return
// 	}
// 	err := rc.RoleService.DeleteRole(roleInput.RoleID)
// 	if err != nil {
// 		if err.Error() == "role not found" {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete role!"})
// 		}
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Role successfully deleted"})
// }
