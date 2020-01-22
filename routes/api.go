package routes

import (
	"github.com/gin-gonic/gin"
	"go-framework/app/Http/Controllers"
)

func Load(router *gin.Engine) {
	router.GET("/", Controllers.HomeController.Index)
	user := router.Group("/users")
	{
		user.POST("", Controllers.UserController.Store)
	}
}