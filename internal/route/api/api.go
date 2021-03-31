package api

import (
	"github.com/gin-gonic/gin"
	"go-framework/core/http"
	route "go-framework/internal/api/v1"
)

func LoadApi(router *gin.Engine) {
	api := router.Group("/api", http.Middleware.Api...)
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/users")
			{
				user.POST("", route.UserStore)
				user.GET("", route.UserList)
			}
		}
	}
}
