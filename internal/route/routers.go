package routes

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-framework/app"
	"go-framework/core/lgo"
	"go-framework/internal/middleware"
	"go-framework/internal/route/api"
	"go-framework/internal/route/swag"
	"io"
	"os"
	"path"
)

var engine *gin.Engine

type Options struct {
	writer    io.Writer
	errWriter io.Writer
}

type Option func(opts *Options)

func WithWriter(w io.Writer) Option {
	return func(opts *Options) {
		opts.writer = w
	}
}

func WithErrWriter(w io.Writer) Option {
	return func(opts *Options) {
		opts.errWriter = w
	}
}

func NewRouter(opts ...Option) *gin.Engine {
	options := Options{
		writer:    os.Stdout,
		errWriter: os.Stderr,
	}

	for _, opt := range opts {
		opt(&options)
	}

	if options.writer != nil {
		gin.DefaultWriter = options.writer
	}

	if options.errWriter != nil {
		gin.DefaultErrorWriter = options.errWriter
	}

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
