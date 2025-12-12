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

type GetProductsV1Param struct {
	Search string `form:"search" binding:"required,min=3,max=50,search"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Email  string `form:"email" binding:"omitempty,email"`
	Date   string `form:"date" binding:"omitempty,datetime=2006-01-02"`
}

type ProductImage struct {
	ImageName string `json:"image_name" binding:"required"`
	ImageLink string `json:"image_link" binding:"required,file_ext=jpg png gif"`
}

type ProductAttribute struct {
	AttributeName  string `json:"attribute_name" binding:"required"`
	AttributeValue string `json:"attribute_value" binding:"required"`
}

type ProductInfo struct {
	InfoKey   string `json:"info_key" binding:"required"`
	InfoValue string `json:"info_value" binding:"required"`
}

type PostProductsV1Param struct {
	Name             string                 `json:"name" binding:"required,min=3,max=100"`
	Price            int                    `json:"price" binding:"required,min_int=100000"`
	Display          *bool                  `json:"display" binding:"omitempty"`
	ProductImage     ProductImage           `json:"product_image" binding:"required"`
	Tags             []string               `json:"tags" binding:"required,gt=3,lt=5"`
	ProductAttribute []ProductAttribute     `json:"product_attribute" binding:"required,gt=0,dive"`
	ProductInfo      map[string]ProductInfo `json:"product_info" binding:"required,gt=0,dive"`
	ProductMetadata  map[string]any         `json:"product_metadata" binding:"omitempty"`
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

func (u *ProductHandler) PostProductsV1(ctx *gin.Context) {
	var params PostProductsV1Param
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	if params.Display == nil {
		defaultDisplay := true
		params.Display = &defaultDisplay
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":           "Create product (V1)",
		"name":              params.Name,
		"price":             params.Price,
		"display":           params.Display,
		"product_image":     params.ProductImage,
		"tags":              params.Tags,
		"product_attribute": params.ProductAttribute,
		"product_info":      params.ProductInfo,
		"product_metadata":  params.ProductMetadata,
	})
}
