package bootstrap

import (
	"go-framework/app"
	"go-framework/conf"
	"go-framework/db"
	"go-framework/db/migrations"
	"go-framework/db/redis"
	"go-framework/internal/cron"
	"go-framework/internal/middleware"
	routes "go-framework/internal/route"
	"go-framework/internal/validator"
	logInit "go-framework/pkg/glog"
	"go-framework/storage"
)

func SetInTest() {
	app.InTest = true
}

func SetInCommand() {
	app.InConsole = true
}

// 应用启动入口
func Bootstrap() {
	conf.Init()

	logInit.Init()
	storage.Init(app.StoragePath)
	// 注册中间件
	middleware.Init()

	if !app.RunningInConsole() {
		// 注册路由
		routes.InitRouter()

		validator.Init()
	}

	redis.Init()

	_ = db.Init()

	migrate()

	if !app.RunningInConsole() {
		// start cron
		cron.Register()
		cron.Start()
	}

	app.IsBootstrap = true
}

func Destroy() {
	db.Close()
	logInit.Close()
	redis.Close()
}

func migrate() {
	db.Conn.AutoMigrate(&migrations.Migration{})
}
