package controllers

import (
	"company_name_eats/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

func (c *AuthController) Register(ctx *gin.Context) {
	var input struct {
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
		SecretKey string `json:"secret_key"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.AuthService.Register(input.Email, input.Password, input.SecretKey)
	if err != nil {
		log.Println("Error registering user:", err)
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}
func (ac *AuthController) Login(c *gin.Context) {
	// Login uchun kerakli malumotlarni olish
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Foydalanuvchi login qilish
	token, err := ac.AuthService.Login(loginData.Email, loginData.Password)
	if err != nil {
		// Xato bo'lsa
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Tokenni muvaffaqiyatli yaratgan taqdirda
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
