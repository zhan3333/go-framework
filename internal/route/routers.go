package routes

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-framework/app"
	"go-framework/core/http"
	storage2 "go-framework/core/storage"
	"go-framework/internal/route/api"
	"go-framework/internal/route/swag"
	"io"
	"os"
	"path"
)

var engine *gin.Engine

func NewRouter() *gin.Engine {
	f, _ := os.Create(storage2.Storage.FullPath("logs/route.log"))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	engine = gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
	// 加载默认中间件
	engine.Use(http.Middleware.Def...)
	pprof.Register(engine)
	loadRoutes()
	return engine
}

// 新增加的路由文件需要在这里进行加载
func loadRoutes() {
	engine.Static("public", path.Join(app.StoragePath, "app/public"))
	api.LoadApi(engine)
	swag.LoadSwag(engine)
}
