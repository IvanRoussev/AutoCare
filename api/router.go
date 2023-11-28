package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()

	// User Login | Anyone can access this route no auth needed
	router.POST("/users/login", server.loginUser)

	// Create User
	router.POST("/users", server.createUser)

	// Auth Middleware protected routes below
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// Owners Routes
	authRoutes.GET("/users/:id", server.getUserByID)
	authRoutes.GET("/users", server.getlistUsers)
	authRoutes.DELETE("/users/:id", server.deleteUserByID)

	// Cars Routes
	authRoutes.POST("/cars", server.createCar)
	authRoutes.GET("/cars/vin/:vin", server.getCarByVIN)
	authRoutes.GET("/cars", server.getListCars)
	authRoutes.GET("/cars/users/:username", server.getListCarsByUsername)
	authRoutes.DELETE("/cars/vin/:vin", server.deleteCarByVIN)

	// Maintenance Routes
	authRoutes.POST("/maintenances", server.createMaintenance)
	authRoutes.GET("/maintenances/:car_vin", server.getListMaintenanceByVIN)
	server.router = router
}
