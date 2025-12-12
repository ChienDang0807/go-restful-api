package v1handler

import (
	"log"
	"net/http"

	"github.com/ChienDang0807/go-restful-api-gin/utils"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

type GetCategoryByCategoryV1Param struct {
	Category string `uri:"category" binding:"oneof=php python golang"`
}

type PostCategoriesV1Param struct {
	Name   string `form:"name" binding:"required"`
	Status string `form:"status" biding:"required,oneof=1 2"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	var params GetCategoryByCategoryV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	log.Println("Into GetCategoryByCategoryV1")

	value, exist := ctx.Get("username")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Missing username"})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Get category by category V1",
		"course":   params.Category,
		"username": value,
	})
}
