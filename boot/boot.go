package boot

import (
	"github.com/zhan3333/gdb"
	"github.com/zhan3333/glog"
	"go-framework/app"
	"go-framework/conf"
	"go-framework/internal/cron"
	"go-framework/internal/middleware"
	routes "go-framework/internal/route"
	"go-framework/internal/validator"
	"go-framework/migrate_file"
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

	glog.DefLogChannel = conf.Logging.Default
	glog.LogConfigs = conf.Logging.Channels
	glog.LoadChannels()

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
	glog.Def().Println("boot success")
}

func Destroy() {
	glog.Close()
	redis.Close()
}
