package api

import (
	"database/sql"
	db "github.com/IvanRoussev/autocare/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createMaintenanceRequest struct {
	CarVin          string `json:"car_vin"`
	MaintenanceType string `json:"maintenance_type"`
	Mileage         int32  `json:"mileage"`
}

func (server *Server) createMaintenance(ctx *gin.Context) {
	var req createMaintenanceRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateMaintenanceParams{
		CarVin:          req.CarVin,
		MaintenanceType: req.MaintenanceType,
		Mileage:         req.Mileage,
	}

	maintenance, err := server.store.CreateMaintenance(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, maintenance)

}

type getListMaintenancesByVINRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

type listMaintenancesByVinRequest struct {
	Vin string `uri:"car_vin" binding:"required,min=1"`
}

func (server *Server) getListMaintenanceByVIN(ctx *gin.Context) {
	var req getListMaintenancesByVINRequest
	var vinReq listMaintenancesByVinRequest

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
