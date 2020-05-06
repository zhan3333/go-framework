package api

import (
	"github.com/gin-gonic/gin"
	"go-framework/app/http/controllers/user_controller"
)

func LoadApi(router *gin.Engine) {
	api := router.Group("/api")
	{
		user := api.Group("/users")
		{
			user.POST("", user_controller.UserController.Store)
			user.GET("", user_controller.UserController.List)
		}
	}
}
