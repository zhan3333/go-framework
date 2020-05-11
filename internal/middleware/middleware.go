package middleware

import "github.com/gin-gonic/gin"

type middleware struct {
	Def []gin.HandlerFunc
	Api []gin.HandlerFunc
}

var Middleware middleware

func Init() {
	Middleware.Def = []gin.HandlerFunc{
		Logger(),
	}
	Middleware.Api = []gin.HandlerFunc{
		ApiTestMiddleware(),
	}
}
