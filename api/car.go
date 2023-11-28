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

type createCarRequest struct {
	Vin      string `json:"vin"`
	Username string `json:"username"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Year     int32  `json:"year"`
}

func (server *Server) createCar(ctx *gin.Context) {
	var req createCarRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateCarParams{
		Vin:      req.Vin,
		Username: authPayload.Username,
		Make:     req.Make,
		Model:    req.Model,
		Year:     req.Year,
	}

	car, err := server.store.CreateCar(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, car)
}

type getCarByVinRequest struct {
	Vin string `uri:"vin" binding:"required,min=1"`
}

func (server *Server) getCarByVIN(ctx *gin.Context) {
	var req getCarByVinRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	car, err := server.store.GetCarByVIN(ctx, req.Vin)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if car.Username != authPayload.Username {
		err := errors.New("car does not belong to authenticated user")
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, car)
}

type listCarsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getListCars(ctx *gin.Context) {
	var req listCarsRequest

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCarsParams{
		Limit:  req.PageSize,
		Offset: req.PageID,
	}

	cars, err := server.store.ListCars(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, cars)

}

type getListCarsByOwnerIDRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

type listCarsByUsernameRequest struct {
	Username string `uri:"username" binding:"required,min=1"`
}

func (server *Server) getListCarsByUsername(ctx *gin.Context) {
	var req getListCarsByOwnerIDRequest
	var ownerIDReq listCarsByUsernameRequest

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = ctx.ShouldBindUri(&ownerIDReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListCarsByUsernameParams{
		Username: authPayload.Username,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
	}

	cars, err := server.store.ListCarsByUsername(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, cars)
}

func (server *Server) deleteCarByVIN(ctx *gin.Context) {
	var req getCarByVinRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	car, err := server.store.GetCarByVIN(ctx, req.Vin)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if car.Username != authPayload.Username {
		errMessage := fmt.Sprintf("can't delete car VIN: %s | Car does not belong to authenticated user: %s", car.Vin, authPayload.Username)
		err := errors.New(errMessage)
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	err = server.store.DeleteCarByVIN(ctx, req.Vin)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
