package api

import (
	db "github.com/IvanRoussev/autocare/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server servers HTTP requests for our Auto Care service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("maintenance_type", validMaintenanceType)
	}

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
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
