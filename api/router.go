package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (server *Server) setupRouter() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"POST", "PUT", "PATCH", "GET", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Bearer", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Anyone can access these routes no auth needed
	// User Login | Anyone can access this route no auth needed
	router.POST("/users/login", server.loginUser)

	// Create User
	router.POST("/users", server.createUser)

	// Auth Middleware protected routes below
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// Owners Routes
	authRoutes.GET("/users/id/:id", server.getUserByID)
	authRoutes.GET("/users/:username", server.getUserByUsername)
	authRoutes.DELETE("/users/:username", server.deleteUserByUsername)

	// Cars Routes
	authRoutes.POST("/cars", server.createCar)
	authRoutes.GET("/cars/vin/:vin", server.getCarByVIN)
	authRoutes.GET("/cars/users/:username", server.getListCarsByUsername)
	authRoutes.DELETE("/cars/vin/:vin", server.deleteCarByVIN)

	// Maintenance Routes
	authRoutes.POST("/maintenances", server.createMaintenance)
	authRoutes.GET("/maintenances/:vin", server.getListMaintenanceByVIN)
	authRoutes.DELETE("/maintenances/:vin", server.deleteMaintenanceByVIN)
	authRoutes.DELETE("/maintenances/id/:id", server.deleteMaintenanceByID)
	server.router = router
}
