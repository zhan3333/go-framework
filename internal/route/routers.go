package routes

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-framework/app"
	"go-framework/core/lgo"
	"go-framework/internal/middleware"
	"go-framework/internal/route/api"
	"go-framework/internal/route/swag"
	"go-framework/pkg/glog"
	"path"
)

var engine *gin.Engine

func NewRouter() *gin.Engine {
	gin.DefaultWriter = glog.Sys.Writer()

	engine = gin.New()

	engine.Use(lgo.WithContext())
	engine.Use(gin.Recovery(), gin.Logger())
	engine.Use(middleware.Logger())

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
