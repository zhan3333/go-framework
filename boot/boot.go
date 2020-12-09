package boot

import (
	"fmt"
	"github.com/zhan3333/gdb/v2"
	"github.com/zhan3333/glog"
	"github.com/zhan3333/go-migrate/v2"
	"github.com/zhan3333/gredis"
	"go-framework/app"
	"go-framework/conf"
	"go-framework/internal/cron"
	"go-framework/internal/middleware"
	routes "go-framework/internal/route"
	"go-framework/internal/validator"
	_ "go-framework/migrate_file"
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

	glog.DefLogChannel = conf.Logging.Default
	glog.LogConfigs = conf.Logging.Channels
	glog.LoadChannels()

	gdb.ConnConfigs = conf.Database.MySQL
	if err = gdb.InitAll(); err != nil {
		glog.Def().Panicf("init gdb module failed: %+v", err)
	}

	gredis.Configs = conf.Database.Redis

	storage.Init(app.StoragePath)

	middleware.Init()

	if !app.RunningInConsole() {
		// 注册路由
		routes.InitRouter()

		validator.Init()
	}

	// load migrate files
	migrate.DB = gdb.Def()
	if err := migrate.InitMigrationTable(); err != nil {
		panic(fmt.Sprintf("migrate.InitMigrationTable() failed: %+v", err))
	}

	if !app.RunningInConsole() {
		// start cron
		cron.Register()
		cron.Start()
	}

	app.Booted = true
	glog.Def().Println("boot success")
}

func Destroy() {
	glog.Close()
	gdb.Close()
	gredis.Close()
}
