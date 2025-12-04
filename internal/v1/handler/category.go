package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

type GetCategoryByCategoryV1Param struct {
	Category string `uri:"category" binding:"oneof=php python golang"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get category by category V1",
	})
}
