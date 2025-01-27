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
func (s *AuthService) Register(email string, password string, secretKey string) (*models.User, error) {
	var existingUser models.User
	if err := s.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("email %s already taken", email)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	isAdmin := false
	adminSecretKey := os.Getenv("ADMIN_SECRET_KEY")
	if secretKey == adminSecretKey {
		isAdmin = true
	}
	var role models.Role
	if isAdmin {
		if err := s.DB.Where("name = ?", "admin").First(&role).Error; err != nil {
			role = models.Role{
				ID:   uuid.New(),
				Name: "admin",
			}
			if result := s.DB.Create(&role); result.Error != nil {
				return nil, fmt.Errorf("failed to create admin role: %v", result.Error)
			}
		}
	} else {
		if err := s.DB.Where("name = ?", "user").First(&role).Error; err != nil {
			role = models.Role{
				ID:   uuid.New(),
				Name: "user",
			}
			if result := s.DB.Create(&role); result.Error != nil {
				return nil, fmt.Errorf("failed to create user role: %v", result.Error)
			}
		}
	}
	user := &models.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
		RoleID:   &role.ID,
	}
	if result := s.DB.Create(&user); result.Error != nil {
		return nil, fmt.Errorf("failed to create user: %v", result.Error)
	}

	return user, nil
}

func (s *AuthService) Login(email string, password string) (string, error) {
	var user models.User
	result := s.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return "", fmt.Errorf("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}
	token, err := s.generateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (s *AuthService) generateJWT(userID uuid.UUID, email string, roleID *uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"email":   email,
		"role_id": roleID,
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
func (s *AuthService) CourierLogin(username string, password string) (string, error) {
	var courier models.Courier
	result := s.DB.Where("username = ?", username).First(&courier)
	if result.Error != nil {
		return "", fmt.Errorf("courier not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(courier.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}
	token, err := s.generateCourierJWT(courier.ID, courier.Username, courier.Fullname)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (s *AuthService) generateCourierJWT(courierID uuid.UUID, username string, full_name string) (string, error) {
	claims := jwt.MapClaims{
		"courier_id": courierID,
		"username":   username,
		"full_name":  full_name,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
		"iat":        time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("could not sign the token: %v", err)
	}

	return signedToken, nil
}
