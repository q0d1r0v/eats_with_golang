package controllers

import (
	"company_name_eats/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (uc *FoodController) CreateFood(c *gin.Context) {
	var foodInput struct {
		CategoryID uuid.UUID `json:"category_id" binding:"required"`
		Name       string    `json:"name" binding:"required"`
		Price      float64   `json:"price" binding:"required"`
	}
	if err := c.ShouldBindJSON(&foodInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category_id name and price is required"})
		return
	}
	food, err := uc.FoodService.CreateFood(foodInput.CategoryID, foodInput.Name, foodInput.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create food!"})
		return
	}
	c.JSON(http.StatusCreated, food)
}
func (uc *FoodController) UpdateFood(c *gin.Context) {
	var foodInput struct {
		FoodID     uuid.UUID `json:"food_id" binding:"required"`
		CategoryID uuid.UUID `json:"category_id"`
		Name       string    `json:"name"`
		Price      float64   `json:"price"`
	}
	if err := c.ShouldBindJSON(&foodInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food_id is required"})
		return
	}
	food, err := uc.FoodService.UpdateFood(foodInput.FoodID, &foodInput.CategoryID, &foodInput.Name, &foodInput.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food!"})
		return
	}
	c.JSON(http.StatusOK, food)
}

func (uc *FoodController) DeleteFood(c *gin.Context) {
	var foodInput struct {
		FoodID uuid.UUID `json:"food_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&foodInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food_id is required"})
		return
	}
	err := uc.FoodService.DeleteFood(foodInput.FoodID)
	if err != nil {
		if err.Error() == "food not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete food!"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Food successfully deleted"})
}
