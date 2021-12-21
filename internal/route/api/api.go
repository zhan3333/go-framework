package api

import (
	"github.com/gin-gonic/gin"

	"go-framework/core/lgo"
	auth2 "go-framework/internal/api/v1/auth"
)

func LoadApi(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			gAuth := v1.Group("/auth")
			{
				gAuth.POST("register", lgo.Route(auth2.Register))
				gAuth.POST("login", lgo.Route(auth2.Login))
			}
		}
	}
}
