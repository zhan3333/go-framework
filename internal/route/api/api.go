package api

import (
	"github.com/gin-gonic/gin"
	"go-framework/internal/controller/user_controller"
	"go-framework/internal/middleware"
)

func LoadApi(router *gin.Engine) {
	api := router.Group("/api", middleware.Middleware.Api...)
	{
		user := api.Group("/users")
		{
			user.POST("", user_controller.Store)
			user.GET("", user_controller.List)
		}
	}
}
