package api

import (
	"github.com/gin-gonic/gin"

	"go-framework/internal/api/v1/auth"
	"go-framework/internal/middleware"
	"go-framework/pkg/lgo"
)

func LoadApi(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			authGroup := v1.Group("/auth")
			{
				authGroup.POST("register", lgo.Controller(auth.Register))
				authGroup.POST("login", lgo.Controller(auth.Login))
			}

			v1.GET("/me", middleware.Auth(), lgo.Controller(auth.Me))
		}
	}
}
