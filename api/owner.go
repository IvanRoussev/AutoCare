package api

//
//import (
//	"database/sql"
//	db "github.com/IvanRoussev/autocare/db/sqlc"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//type createUserRequest struct {
//	FullName string `json:"full_name" binding:"required"`
//	Username string `json:"username" binding:"required"`
//	password string `json:"password"`
//	Country  string `json:"country" binding:"required"`
//}
//
//func (server *Server) createOwner(ctx *gin.Context) {
//	var req createOwnerRequest
//	err := ctx.ShouldBindJSON(&req)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	arg := db.CreateOwnerParams{
//		FirstName: req.FirstName,
//		LastName:  req.LastName,
//		Country:   req.Country,
//	}
//
//	owner, err := server.store.CreateOwner(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	ctx.JSON(http.StatusOK, owner)
//}
//
//type getOwnerByIDRequest struct {
//	ID int64 `uri:"id" binding:"required,min=1"`
//}
//
//func (server *Server) getOwnerByID(ctx *gin.Context) {
//	var req getOwnerByIDRequest
//	err := ctx.ShouldBindUri(&req)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	owner, err := server.store.GetOwnerByID(ctx, req.ID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			ctx.JSON(http.StatusNotFound, errorResponse(err))
//			return
//		}
//
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	ctx.JSON(http.StatusOK, owner)
//}
//
//type listOwnerRequest struct {
//	PageID   int32 `form:"page_id" binding:"required,min=1"`
//	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
//}
//
//func (server *Server) getlistOwners(ctx *gin.Context) {
//	var req listOwnerRequest
//
//	err := ctx.ShouldBindQuery(&req)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	arg := db.ListOwnersParams{
//		Limit:  req.PageSize,
//		Offset: (req.PageID - 1) * req.PageSize,
//	}
//
//	owners, err := server.store.ListOwners(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	ctx.JSON(http.StatusOK, owners)
//}
//
//type deleteOwnerByIDRequest struct {
//	ID int64 `uri:"id" binding:"required,min=1"`
//}
//
//func (server *Server) deleteOwnerByID(ctx *gin.Context) {
//	var req deleteOwnerByIDRequest
//	err := ctx.ShouldBindUri(&req)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	err = server.store.DeleteOwnerByID(ctx, req.ID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			ctx.JSON(http.StatusNotFound, errorResponse(err))
//			return
//		}
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	ctx.JSON(http.StatusNoContent, gin.H{"message": "Owner deleted successfully"})
//}
