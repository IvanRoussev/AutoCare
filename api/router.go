package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()

	// User Login
	router.POST("/users/login", server.loginUser)

	// Owners Routes
	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUserByID)
	router.GET("/users", server.getlistUsers)
	router.DELETE("/users/:id", server.deleteUserByID)

	// Cars Routes
	router.POST("/cars", server.createCar)
	router.GET("/cars/vin/:vin", server.getCarByVIN)
	router.GET("/cars", server.getListCars)
	router.GET("/cars/owner/:owner_id", server.getListCarsByOwnerID)
	router.DELETE("/cars/:vin", server.deleteCarByVIN)

	// Maintenance Routes
	router.POST("/maintenances", server.createMaintenance)
	router.GET("/maintenances/:car_vin", server.getListMaintenanceByVIN)
	server.router = router
}
