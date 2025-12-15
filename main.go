package main

import (
	v1handler "github.com/ChienDang0807/go-restful-api-gin/internal/v1/handler"
	"github.com/ChienDang0807/go-restful-api-gin/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	if err := utils.RegisterValidators(); err != nil {
		panic(err)
	}

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			user.GET("/", userHandlerV1.GetUsersV1)
			user.GET("/admin/:uuid", userHandlerV1.GetUserByUUID)
			user.GET("/:id", userHandlerV1.GetUserByIdV2)
		}
		category := v1.Group("/categories")
		{
			categoryHandlerV1 := v1handler.NewCategoryHandler()
			category.GET("/:category", categoryHandlerV1.GetCategoryByCategoryV1)
			category.POST("/", categoryHandlerV1.PostCategoriesV1)
		}

		news := v1.Group("/news")
		{
			newsHandlerV1 := v1handler.NewNewsHandler()
			news.GET("/", newsHandlerV1.GetNewsV1)
			// news.GET("/:slug", middleware.SimpleMiddleware(), newsHandlerV1.GetNewsV1)
			news.POST("/", newsHandlerV1.PostNewsV1)
			news.POST("/upload-file", newsHandlerV1.PostUploadFileNewsV1)
			news.POST("/upload-multiple-file", newsHandlerV1.PostUploadMultipleFileNewsV1)
		}
	}

	r.Run(":8080")
}
