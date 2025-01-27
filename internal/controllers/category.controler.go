package controllers

import (
	"company_name_eats/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryController struct {
	CategoryService *services.CategoryService
}

func (uc *CategoryController) LoadCategories(c *gin.Context) {
	categories, err := uc.CategoryService.LoadCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories!"})
		return
	}
	c.JSON(http.StatusOK, categories)
}
func (uc *CategoryController) CreateCategory(c *gin.Context) {
	var categoryInput struct {
		BranchID uuid.UUID `json:"branch_id" binding:"required"`
		Name     string    `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&categoryInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name and branch_id is required"})
		return
	}
	category, err := uc.CategoryService.CreateCategory(categoryInput.BranchID, categoryInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category!"})
		return
	}
	c.JSON(http.StatusCreated, category)
}

func (uc *CategoryController) UpdateCategory(c *gin.Context) {
	var categoryInput struct {
		CatgoryID uuid.UUID `json:"category_id" binding:"required"`
		BranchID  uuid.UUID `json:"branch_id" binding:"required"`
		Name      string    `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&categoryInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category category_id, name and branch_id is required"})
		return
	}
	category, err := uc.CategoryService.UpdateCategory(categoryInput.CatgoryID, categoryInput.BranchID, categoryInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category!"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (uc *CategoryController) DeleteCategory(c *gin.Context) {
	var categoryInput struct {
		CatgoryID uuid.UUID `json:"category_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&categoryInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category_id is required"})
		return
	}
	err := uc.CategoryService.DeleteCategory(categoryInput.CatgoryID)
	if err != nil {
		if err.Error() == "category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category!"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category successfully deleted"})
}
