package api

import (
	"github.com/gin-gonic/gin"
	"go-framework/core/http"
	user2 "go-framework/internal/api/v1/user"
)

func LoadApi(router *gin.Engine) {
	api := router.Group("/api", http.Middleware.Api...)
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("register", user2.Register)
			}
			user := v1.Group("/users")
			{
				user.GET("", user2.List)
			}
		}
	}
}
