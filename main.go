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

		news := v1.Group("/news")
		{
			newHandlerV1 := v1handler.NewNewsHandler()
			news.GET("/", newHandlerV1.GetNewsV1)
			news.GET("/:slug", newHandlerV1.GetNewsV1)
		}
	}

	r.Run(":8080")
}
