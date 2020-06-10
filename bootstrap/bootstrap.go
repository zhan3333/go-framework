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
	var err error
	conf.Init()

	logInit.Init()

	storage.Init(app.StoragePath)

	middleware.Init()

	if !app.RunningInConsole() {
		// 注册路由
		routes.InitRouter()

		validator.Init()
	}

	_, err = redis.InitDef()
	if err != nil {
		panic(err)
	}

	// 连接默认数据库
	_, err = db.InitDef()
	if err != nil {
		panic(err)
	}

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
	db.Def().AutoMigrate(&migrations.Migration{})
}
