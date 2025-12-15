package v1handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ChienDang0807/go-restful-api-gin/utils"
	"github.com/gin-gonic/gin"
)

type NewHandler struct {
}

func NewNewsHandler() *NewHandler {
	return &NewHandler{}
}

type PostNewsV1Param struct {
	Title  string `form:"title" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
}

func (n *NewHandler) GetNewsV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if slug == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get new V1",
			"slug":    "No News",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get new V1",
			"slug":    slug,
		})
	}
}

func (n *NewHandler) PostNewsV1(ctx *gin.Context) {
	var params PostNewsV1Param
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	err = os.MkdirAll("/uploads", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create upload folder"})
		return
	}

	dst := fmt.Sprintf("./upload/%s", filepath.Base(image.Filename))
	if err := ctx.SaveUploadedFile(image, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save file"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post news (V1)",
		"title":   params.Title,
		"status":  params.Status,
		"image":   image.Filename,
		"path":    dst,
	})
}

func (n *NewHandler) PostUploadFileNewsV1(ctx *gin.Context) {

}

func (n *NewHandler) PostUploadMultipleFileNewsV1(ctx *gin.Context) {

}
