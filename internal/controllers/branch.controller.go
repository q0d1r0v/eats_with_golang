package controllers

import (
	"company_name_eats/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BranchController struct {
	BranchService *services.BranchService
}

func (uc *BranchController) LoadBranches(c *gin.Context) {
	branches, err := uc.BranchService.LoadBranches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch branches!"})
		return
	}
	c.JSON(http.StatusOK, branches)
}
func (uc *BranchController) CreateBranch(c *gin.Context) {
	var branchInput struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&branchInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch name is required"})
		return
	}
	branch, err := uc.BranchService.CreateBranch(branchInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create branch!"})
		return
	}
	c.JSON(http.StatusCreated, branch)
}
func (uc *BranchController) UpdateBranch(c *gin.Context) {
	var branchInput struct {
		ID   uuid.UUID `json:"branch_id" binding:"required"`
		Name string    `json:"branch_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&branchInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch_id and branch_name is required"})
		return
	}
	branch, err := uc.BranchService.UpdateBranchNameByID(branchInput.ID, branchInput.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update branch!"})
		return
	}
	c.JSON(http.StatusOK, branch)
}
func (uc *BranchController) DeleteBranch(c *gin.Context) {
	var branchInput struct {
		ID uuid.UUID `json:"branch_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&branchInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch_id is required"})
		return
	}
	err := uc.BranchService.DeleteBranchByID(branchInput.ID)
	if err != nil {
		if err.Error() == "branch not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete branch!"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Branch successfully deleted"})
}
