package api

import (
	"github.com/gin-gonic/gin"
	"go-framework/core/http"
	auth2 "go-framework/internal/api/v1/auth"
	user2 "go-framework/internal/api/v1/user"
	"go-framework/internal/middleware"
)

func LoadApi(router *gin.Engine) {
	api := router.Group("/api", http.Middleware.Api...)
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("register", auth2.Register)
				auth.POST("login", auth2.Login)
			}
			user := v1.Group("/users", middleware.Auth())
			{
				user.GET("", user2.List)
			}
		}
	}
}
