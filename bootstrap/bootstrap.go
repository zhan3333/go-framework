package bootstrap

import (
	"go-framework/app"
	"go-framework/db"
	"go-framework/db/migrations"
	"go-framework/db/redis"
	"go-framework/internal/middleware"
	"go-framework/internal/route"
	"go-framework/internal/validator"
	"go-framework/log"
)

func SetInTest() {
	app.App.InTest = true
}

func SetInCommand() {
	app.App.InCommand = true
}

// 应用启动入口
func Bootstrap() {
	//migrate()

	app.Init()

	log.Init()

	// 注册中间件
	middleware.Init()

	// 注册路由
	routes.InitRouter()

	validator.Init()

	redis.Init()

	db.Init()

	migrate()

	app.App.IsBootstrap = true
}

func Destroy() {
	db.Close()
	log.Close()
	redis.Close()
}

func migrate() {
	db.Conn.AutoMigrate(&migrations.Migration{})
}
