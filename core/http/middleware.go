package http

import (
	"github.com/gin-gonic/gin"
	"go-framework/internal/middleware"
)

var Middleware struct {
	Def []gin.HandlerFunc
	Api []gin.HandlerFunc
}

func Init() {
	Middleware.Def = []gin.HandlerFunc{
		middleware.Logger(),
	}
	Middleware.Api = []gin.HandlerFunc{
		middleware.WithResp(),
	}
}
