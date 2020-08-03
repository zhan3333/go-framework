package boot

import (
	log "github.com/sirupsen/logrus"
	"github.com/zhan3333/gdb"
	"go-framework/app"
	"go-framework/conf"
	"go-framework/internal/cron"
	"go-framework/internal/middleware"
	routes "go-framework/internal/route"
	"go-framework/internal/validator"
	"go-framework/migrate_file"
	logInit "go-framework/pkg/glog"
	"go-framework/pkg/redis"
	"go-framework/storage"
)

func SetInTest() {
	app.InTest = true
}

func SetInCommand() {
	app.InConsole = true
}

// 应用启动入口
func Boot() {
	var err error
	conf.Init()
	gdb.ConnConfigs = conf.Database.MySQL

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

	// load migrate files
	migrate_file.Init()

	if !app.RunningInConsole() {
		// start cron
		cron.Register()
		cron.Start()
	}

	app.IsBootstrap = true
	log.Println("boot success")
}

func Destroy() {
	logInit.Close()
	redis.Close()
}
