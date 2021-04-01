package boot

import (
	"fmt"
	"github.com/zhan3333/gdb/v2"
	"github.com/zhan3333/glog"
	"github.com/zhan3333/go-migrate/v2"
	"github.com/zhan3333/gredis"
	"go-framework/app"
	"go-framework/conf"
	"go-framework/core/http"
	storage2 "go-framework/core/storage"
	"go-framework/internal/cron"
	routes "go-framework/internal/route"
	"go-framework/internal/validator"
	_ "go-framework/migrate_file"
	"time"
)

// 框架启动

func SetInTest() {
	app.InTest = true
}

func SetInCommand() {
	app.InConsole = true
}

// 应用启动入口
func Boot() {
	bootStart()

	bootLog()
	bootDB()
	bootStorage()
	bootHTTP()
	bootSchedule()

	app.Booted = true
	glog.Def().Println("boot success")
}

func Destroy() {
	glog.Close()
	gdb.Close()
	gredis.Close()
}

func bootStart() {
	datetime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] boot: %s\n", datetime, "app boot start")
}

func logBootInfo(info string) {
	datetime := time.Now().Format("2006-01-02 15:04:05")
	glog.Def().Infof("[%s] boot: %s", datetime, info)
	fmt.Printf("[%s] boot: %s\n", datetime, info)
}

func logBootPanic(msg string, err error) {
	datetime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] boot: %s: %+v\n", datetime, msg, err)
	glog.Def().Panicf("[%s] boot: %s: %+v", datetime, msg, err)
}

func bootLog() {
	glog.DefLogChannel = conf.Logging.Default
	glog.LogConfigs = conf.Logging.Channels
	glog.LoadChannels()
	logBootInfo("log module init")
}

func bootDB() {
	gdb.ConnConfigs = conf.Database.MySQL
	conf.InitDBLogMode()
	if err := gdb.InitAll(); err != nil {
		logBootPanic("init gdb module failed: %+v", err)
	}

	gredis.Configs = conf.Database.Redis

	// load migrate files
	migrate.DB = gdb.Def()
	if err := migrate.InitMigrationTable(); err != nil {
		logBootPanic("migrate.InitMigrationTable() failed", err)
	}
	logBootInfo("database module init")
}

func bootStorage() {
	storage2.Init(app.StoragePath)
	logBootInfo("storage module init")
}

func bootHTTP() {
	http.Init()
	logBootInfo("middleware module init")
	if !app.RunningInConsole() {
		// 注册路由
		app.SetRouter(routes.NewRouter())
		logBootInfo("route module init")

		validator.Init()
		logBootInfo("validator module init")
	}
}

func bootSchedule() {
	if !app.RunningInConsole() {
		cron.Register()
		cron.Start()
		logBootInfo("schedule module start")
	}
}
