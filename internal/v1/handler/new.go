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

	if image.Size > 5<<20 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File too large (5 MB) "})
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

	filename, err := utils.ValidateAndSaveFile(image, "./uploads")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post news (V1)",
		"title":   params.Title,
		"status":  params.Status,
		"image":   image.Filename,
		"path":    "./upload/" + filename,
	})
}

func (n *NewHandler) PostUploadMultipleFileNewsV1(ctx *gin.Context) {
	const publicURL = "http://localhost:8080/images/"
	var params PostNewsV1Param
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart form"})
		return
	}

	images := form.File["images"]
	if len(images) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}

	var successFiles []string
	var failFiles []map[string]string
	for _, image := range images {
		fileName, err := utils.ValidateAndSaveFile(image, "./uploads")
		if err != nil {
			failFiles = append(failFiles, map[string]string{
				"filename": image.Filename,
				"error":    err.Error(),
			})
			continue
		}
		publicImageURL := publicURL + fileName
		successFiles = append(successFiles, publicImageURL)
	}

	resp := gin.H{
		"message":       "Post news (V1)",
		"title":         params.Title,
		"status":        params.Status,
		"success_files": successFiles,
	}

	if len(failFiles) > 0 {
		resp["message"] = "Upload completed with partial errors"
		resp["error_files"] = failFiles
	}

	ctx.JSON(http.StatusOK, resp)
}
