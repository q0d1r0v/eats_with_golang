package services

import (
	"company_name_eats/internal/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// register user
func (s *AuthService) Register(email string, password string) (*models.User, error) {
	// check to unique email
	var existingUser models.User
	if err := s.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("email %s already taken", email)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
	}
	result := s.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// login
func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User
	result := s.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return "", fmt.Errorf("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}
	token, err := s.generateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// JWT yaratish
func (s *AuthService) generateJWT(userID uuid.UUID, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("could not sign the token: %v", err)
	}

	return signedToken, nil
}
