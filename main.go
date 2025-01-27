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
	// userService := &services.UserService{DB: db}
	// userController := &controllers.UserController{UserService: userService}
	authService := &services.AuthService{DB: db}
	authController := &controllers.AuthController{AuthService: authService}
	foodService := &services.FoodService{DB: db}
	foodController := &controllers.FoodController{FoodService: foodService}
	branchService := &services.BranchService{DB: db}
	branchController := &controllers.BranchController{BranchService: branchService}
	categoryService := &services.CategoryService{DB: db}
	categoryController := &controllers.CategoryController{CategoryService: categoryService}
	roleService := &services.RoleService{DB: db}
	roleController := &controllers.RoleController{RoleService: roleService}
	courierService := &services.CourierService{DB: db}
	courierController := &controllers.CourierController{CourierService: courierService}
	orderService := &services.OrderService{DB: db}
	orderController := &controllers.OrderController{OrderService: orderService}

	// groups
	admin := r.Group("/admin/")
	courier := r.Group("/courier/")
	auth_route := r.Group("/auth/")
	api_route := r.Group("/api/")

	// use middleware
	admin.Use(middlewares.JWTAdminMiddleware(db))
	courier.Use(middlewares.JWTCourierMiddleware(db))
	api_route.Use(middlewares.JWTAuthMiddleware())

	// auth routes
	auth_route.POST("/register", authController.Register)
	auth_route.POST("/login", authController.Login)
	auth_route.POST("/courier/login", authController.CourierLogin)

	// food routes
	api_route.GET("/v1/load/foods", foodController.LoadFoods)
	api_route.POST("/v1/order/food", orderController.CreateOrder)

	// courier roures
	courier.GET("/api/v1/load/all/orders", orderController.LoadAllOrders)
	courier.PUT("/api/v1/update/order", orderController.UpdateOrder)

	// admin branch routes
	admin.POST("/api/v1/create/branch", branchController.CreateBranch)
	admin.GET("/api/v1/load/branches", branchController.LoadBranches)
	admin.PUT("/api/v1/update/branch", branchController.UpdateBranch)
	admin.DELETE("/api/v1/delete/branch", branchController.DeleteBranch)

	// admin category routes
	admin.GET("/api/v1/load/categories", categoryController.LoadCategories)
	admin.POST("/api/v1/create/category", categoryController.CreateCategory)
	admin.PUT("/api/v1/update/category", categoryController.UpdateCategory)
	admin.DELETE("/api/v1/delete/category", categoryController.DeleteCategory)

	// admin food routes
	admin.GET("/api/v1/load/foods", foodController.LoadFoods)
	admin.POST("/api/v1/create/food", foodController.CreateFood)
	admin.PUT("/api/v1/update/food", foodController.UpdateFood)
	admin.DELETE("/api/v1/delete/food", foodController.DeleteFood)

	// admin role routes
	admin.GET("/api/v1/load/roles", roleController.LoadRoles)
	admin.POST("/api/v1/create/role", roleController.CreateRole)
	admin.PUT("/api/v1/update/role", roleController.UpdateRole)
	admin.DELETE("/api/v1/delete/role", roleController.DeleteRole)

	// admin courier routes
	admin.GET("/api/v1/load/couriers", courierController.LoadCouriers)
	admin.POST("/api/v1/create/courier", courierController.CreateCourier)
	admin.PUT("/api/v1/update/courier", courierController.UpdateCourier)
	// admin.DELETE("/api/v1/delete/couriers", courierController.DeleteCourier)

	// run server
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
	fmt.Printf("Server is running on port %s\n", port)
}
