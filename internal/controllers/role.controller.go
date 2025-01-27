package controllers

import (
	"company_name_eats/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleController struct {
	RoleService *services.RoleService
}

func (rc *RoleController) LoadRoles(c *gin.Context) {
	roles, err := rc.RoleService.LoadRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch roles"})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (rc *RoleController) CreateRole(c *gin.Context) {
	var roleInput struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	role, err := rc.RoleService.CreateRole(roleInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role!"})
		return
	}
	c.JSON(http.StatusCreated, role)
}

func (uc *RoleController) UpdateRole(c *gin.Context) {
	var roleInput struct {
		RoleID uuid.UUID `json:"food_id" binding:"required"`
		Name   string    `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role_id and name are required"})
		return
	}
	role, err := uc.RoleService.UpdateRole(roleInput.RoleID, &roleInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role!"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func (rc *RoleController) DeleteRole(c *gin.Context) {
	var roleInput struct {
		RoleID uuid.UUID `json:"role_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role_id is required"})
		return
	}
	err := rc.RoleService.DeleteRole(roleInput.RoleID)
	if err != nil {
		if err.Error() == "role not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete role!"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role successfully deleted"})
}
