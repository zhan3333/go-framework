package routes

import (
	"github.com/gin-gonic/gin"
	"go-framework/conf"
	"go-framework/internal/middleware"
	"go-framework/internal/route/api"
	"go-framework/internal/route/swag"
	"go-framework/storage"
	"io"
	"os"
)

var router *gin.Engine

func InitRouter() {
	gin.SetMode(conf.GinModel)
	f, _ := os.Create(storage.Storage.FullPath("logs/route.log"))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router = gin.New()
	router.Use(gin.Recovery(), gin.Logger())
	// 加载默认中间件
	router.Use(middleware.Middleware.Def...)
	loadRoutes()
}

// 新增加的路由文件需要在这里进行加载
func loadRoutes() {
	api.LoadApi(router)
	swag.LoadSwag(router)
}

func GetRouter() *gin.Engine {
	if router == nil {
		InitRouter()
	}
	return router
}
