package api

import (
	"database/sql"
	"errors"
	"fmt"
	db "github.com/IvanRoussev/autocare/db/sqlc"
	"github.com/IvanRoussev/autocare/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createMaintenanceRequest struct {
	CarVin          string `json:"car_vin"`
	MaintenanceType string `json:"maintenance_type" `
	Mileage         int32  `json:"mileage"`
}

type getListMaintenancesByVINRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

type getMaintenancesByVinRequest struct {
	Vin string `uri:"vin" binding:"required,min=1"`
}

type getMaintenanceByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) createMaintenance(ctx *gin.Context) {
	var req createMaintenanceRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMaintenanceParams{
		CarVin:          req.CarVin,
		MaintenanceType: req.MaintenanceType,
		Mileage:         req.Mileage,
	}
	car, err := server.store.GetCarByVIN(ctx, req.CarVin)

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if car.Username != authPayload.Username {
		err := errors.New("car does not belong to authenticated user")
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	maintenance, err := server.store.CreateMaintenance(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, maintenance)

}

func (server *Server) getListMaintenanceByVIN(ctx *gin.Context) {
	var req getListMaintenancesByVINRequest
	var vinReq getMaintenancesByVinRequest

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = ctx.ShouldBindUri(&vinReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetListMaintenancesByVINParams{
		CarVin: vinReq.Vin,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	maintenances, err := server.store.GetListMaintenancesByVIN(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, maintenances)
}

func (server *Server) deleteMaintenanceByVIN(ctx *gin.Context) {
	var req getMaintenancesByVinRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	car, err := server.store.GetCarByVIN(ctx, req.Vin)
	if car.Username != authPayload.Username {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteMaintenanceByVIN(ctx, req.Vin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	msg := fmt.Sprintf("Successfully deleted maintenances with VIN: %v", req.Vin)
	ctx.JSON(http.StatusNoContent, gin.H{"success": msg})
}

func (server *Server) deleteMaintenanceByID(ctx *gin.Context) {
	var req getMaintenanceByIDRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	maintenance, err := server.store.GetMaintenanceByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}
	car, err := server.store.GetCarByVIN(ctx, maintenance.CarVin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if car.Username != authPayload.Username {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	err = server.store.DeleteMaintenanceByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	msg := fmt.Sprintf("Successfully deleted maintenance with ID: %v", req.ID)
	ctx.JSON(http.StatusNoContent, gin.H{"success": msg})
}
