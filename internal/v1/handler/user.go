package v1handler

import (
	"net/http"
	"strconv"

	"github.com/ChienDang0807/go-restful-api-gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
}

type GetUsersByIdV2Param struct {
	ID int `uri:"id" binding:"gt=0" `
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"OK",
	})
}

func (u *UserHandler) GetUserByIdV2 (ctx *gin.Context){
	var params GetUsersByIdV2Param

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by user Id",
		"userId": params.ID,
	})
}

func (u *UserHandler) GetUserByIdV1 (ctx *gin.Context){
	idStr := ctx.Param("id")
	id , err := strconv.Atoi(idStr)
	if err!= nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "Id must be a number",
		})
		return
	}

	if id<= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"ID must be positive",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID",

	})
}


func (u *UserHandler) GetUserByUUID (ctx *gin.Context){
	uuidStr := ctx.Param("uuid")

	_, err := uuid.Parse(uuidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" :"Id must be a valid UUId",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID",

	})
}