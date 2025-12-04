package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewHandler struct {
}

func NewNewsHandler() *NewHandler {
	return &NewHandler{}
}

func (n *NewHandler) GetNewsV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if (slug == "") {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get new V1",
			"slug": "No News",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get new V1",
			"slug":slug,
		})
	}
}
