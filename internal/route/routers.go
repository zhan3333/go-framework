package routes

import (
	"github.com/gin-gonic/gin"
	"go-framework/conf"
	"go-framework/core/http"
	"go-framework/internal/route/api"
	"go-framework/internal/route/swag"
	"go-framework/storage"
	"io"
	"os"
)

var engine *gin.Engine

func NewRouter() *gin.Engine {
	gin.SetMode(conf.GinModel)
	f, _ := os.Create(storage.Storage.FullPath("logs/route.log"))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	engine = gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
	// 加载默认中间件
	engine.Use(http.Middleware.Def...)
	loadRoutes()
	return engine
}

// 新增加的路由文件需要在这里进行加载
func loadRoutes() {
	api.LoadApi(engine)
	swag.LoadSwag(engine)
}
