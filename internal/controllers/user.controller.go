package controllers

import (
	"company_name_eats/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
