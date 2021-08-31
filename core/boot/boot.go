package boot

import (
	"errors"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"go-framework/app"
	"go-framework/conf"
	"go-framework/core/http"
	storage2 "go-framework/core/storage"
	"go-framework/internal/cron"
	routes "go-framework/internal/route"
	"go-framework/internal/validator"
	_ "go-framework/migrate_file"
	"go-framework/pkg/gdb"
	"go-framework/pkg/glog"
	"go-framework/pkg/gredis"
	"go-framework/pkg/migrate"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path"
	"strconv"
	"time"
)

// 框架启动

func SetInTest() {
	app.InTest = true
}

func SetInCommand() {
	app.InConsole = true
}

// Boot 应用启动入口
func Boot() error {
	err := func() error {
		if config, err := loadConfig(); err != nil {
			return fmt.Errorf("load config failed: %w", err)
		} else {
			app.Config = config
		}
		if err := bootLog(); err != nil {
			return err
		}

		logBootInfo("boot start")

		if err := bootDB(); err != nil {
			return err
		}

		logBootInfo("database module init")

		bootStorage()
		bootHTTP()

		if app.Config.Cron.Enable {
			bootSchedule()
		}

		app.Booted = true
		logBootInfo("boot success")
		return nil
	}()
	if err != nil {
		datetime := time.Now().Format("2006-01-02 15:04:05")
		glog.Default.Errorf("[%s] boot failed: %s", datetime, err)
	}
	return nil
}

func Destroy() {
	gdb.Close()
	gredis.Close()
}

func loadConfig() (*conf.Config, error) {
	var defaultConfigFile = "local.toml"
	var configFile *string
	var config = conf.Config{}

	if app.InConsole {
		configFile = &defaultConfigFile
	} else {
		configFile = kingpin.Flag("config", "load config file").Default(defaultConfigFile).String()
		kingpin.Parse()
	}

	if err := configor.Load(&config, path.Join("config", *configFile)); err != nil {
		return nil, err
	}
	return &config, nil
}

func logBootInfo(info string) {
	datetime := time.Now().Format("2006-01-02 15:04:05")
	glog.Default.Infof("[%s] boot: %s", datetime, info)
}

func bootLog() error {
	var ok bool
	var sysLog conf.Log
	var defLog conf.Log
	if sysLog, ok = app.Config.Log["sys"]; !ok {
		sysLog = conf.DefaultLog
	}
	if defLog, ok = app.Config.Log["sys"]; !ok {
		defLog = conf.DefaultLog
	}
	glog.Sys = logrus.New()
	sysLevel, err := logrus.ParseLevel(sysLog.Level)
	if err != nil {
		return err
	}

	glog.Sys.SetLevel(sysLevel)
	switch sysLog.Write {
	case "stderr":
		glog.Sys.SetOutput(os.Stderr)
	case "file":
		if sysLog.FilePath == "" {
			return errors.New("sys log file path no set")
		}
		f, err := os.Create(sysLog.FilePath)
		if err != nil {
			return fmt.Errorf("open sys log file %s failed: %w", sysLog.FilePath, err)
		}
		glog.Sys.SetOutput(f)
	default:
		return fmt.Errorf("no supported write: %s", sysLog.Write)
	}

	glog.Default = logrus.New()
	defLevel, err := logrus.ParseLevel(defLog.Level)
	if err != nil {
		return err
	}

	glog.Default.SetLevel(defLevel)
	switch sysLog.Write {
	case "stderr":
		glog.Sys.SetOutput(os.Stderr)
	case "file":
		if sysLog.FilePath == "" {
			return errors.New("sys log file path no set")
		}
		f, err := os.Create(sysLog.FilePath)
		if err != nil {
			return fmt.Errorf("open sys log file %s failed: %w", sysLog.FilePath, err)
		}
		glog.Sys.SetOutput(f)
	default:
		return fmt.Errorf("no supported write: %s", sysLog.Write)
	}
	logBootInfo("log module init")
	return nil
}

func bootDB() error {
	for k, v := range app.Config.Databases {
		gdb.ConnConfigs[k] = gdb.MySQLConf{
			Host:     v.Host,
			Port:     strconv.Itoa(v.Port),
			Username: v.Username,
			Password: v.Password,
			Database: v.Database,
		}
	}
	if err := gdb.InitAll(); err != nil {
		return fmt.Errorf("init gdb module failed: %w", err)
	}

	for k, v := range app.Config.Redis {
		gredis.Configs[k] = gredis.Conf{
			Host:     v.Host,
			Password: v.Password,
			Port:     v.Port,
			Database: v.Index,
		}
	}

	// load migrate files
	migrate.DB = gdb.Def()
	if err := migrate.InitMigrationTable(); err != nil {
		return fmt.Errorf("migrate.InitMigrationTable() failed %w", err)
	}
	return nil
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
