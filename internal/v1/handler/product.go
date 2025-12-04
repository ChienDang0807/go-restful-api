package v1handler

import (
	"net/http"
	"regexp"

	"github.com/ChienDang0807/go-restful-api-gin/utils"
	"github.com/gin-gonic/gin"
)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.])*$`)

type ProductHandler struct {
}

type GetProductsBySlugV1Param struct {
	Slug string `uri:"slug" biding:"slug"`
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductsv1(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all product (v1)",
		"limit":   limit,
	})
}

func (p *ProductHandler) GetProducBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if !slugRegex.MatchString(slug) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Sai slug",
		})
	}
}

func (p *ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {
	var params GetProductsBySlugV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Get product by Slug",
		"slug":    params.Slug,
	})
}
