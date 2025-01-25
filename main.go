package main

import (
	"company_name_eats/internal/controllers"
	"company_name_eats/internal/middlewares"
	"company_name_eats/internal/models"
	"company_name_eats/internal/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// load
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}

	// load env data
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSslMode := os.Getenv("DB_SSLMODE")
	dbTimezone := os.Getenv("DB_TIMEZONE")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// connection to database
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, dbSslMode, dbTimezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// auto migrate
	if err := db.AutoMigrate(&models.Branch{}, &models.Category{}, &models.Courier{}, &models.Food{}, &models.OrderItem{}, &models.Order{}, &models.Role{}, &models.User{}); err != nil {
		log.Fatal("Error during migration:", err)
	}

	// gin route
	r := gin.Default()

	// use services and controllers
	userService := &services.UserService{DB: db}
	userController := &controllers.UserController{UserService: userService}
	authService := &services.AuthService{DB: db}
	authController := &controllers.AuthController{AuthService: authService}
	foodService := &services.FoodService{DB: db}
	foodController := &controllers.FoodController{FoodService: foodService}

	// groups
	admin := r.Group("/admin/")
	auth_route := r.Group("/auth/")
	api_route := r.Group("/api/")

	// use middleware
	api_route.Use(middlewares.JWTAuthMiddleware())

	// routes
	auth_route.POST("/register", authController.Register)
	auth_route.POST("/login", authController.Login)
	admin.GET("/api/v1/users", userController.GetAllUsers)
	api_route.GET("/v1/load/foods", foodController.LoadFoods)

	// run server
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
	fmt.Printf("Server is running on port %s\n", port)
}
