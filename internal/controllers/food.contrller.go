package controllers

import (
	"company_name_eats/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoodController struct {
	FoodService *services.FoodService
}

func (uc *FoodController) LoadFoods(c *gin.Context) {
	foods, err := uc.FoodService.LoadFoods()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch foods"})
		return
	}
	c.JSON(http.StatusOK, foods)
}
