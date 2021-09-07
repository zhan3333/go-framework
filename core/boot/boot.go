package boot

import (
	"errors"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"go-framework/app"
	"go-framework/conf"
	gdb2 "go-framework/core/gdb"
	glog2 "go-framework/core/glog"
	gredis2 "go-framework/core/gredis"
	migrate2 "go-framework/core/migrate"
	storage2 "go-framework/core/storage"
	"go-framework/internal/cron"
	routes "go-framework/internal/route"
	"go-framework/internal/validator"
	_ "go-framework/migrate_file"
	"os"
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

type Options struct {
	configFile       string
	enableRoutePrint bool
}

type Option func(opts *Options)

func WithConfigFile(configFile string) Option {
	return func(opts *Options) {
		opts.configFile = configFile
	}
}

func WithRoutePrint(enable bool) Option {
	return func(opts *Options) {
		opts.enableRoutePrint = enable
	}
}

// New 应用启动入口
func New(opts ...Option) error {
	options := Options{
		enableRoutePrint: true,
	}
	for _, opt := range opts {
		opt(&options)
	}

	err := func() error {
		if config, err := loadConfig(options.configFile); err != nil {
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
		glog2.Default.Errorf("[%s] boot failed: %s", datetime, err)
	}
	return nil
}

func Destroy() {
	gdb2.Close()
	gredis2.Close()
}

func loadConfig(configFile string) (*conf.Config, error) {
	var config = conf.Config{}
	if err := configor.Load(&config,
		configFile,
		"config/default.toml",
	); err != nil {
		return nil, err
	}
	return &config, nil
}

func logBootInfo(info string) {
	datetime := time.Now().Format("2006-01-02 15:04:05")
	glog2.Default.Infof("[%s] boot: %s", datetime, info)
}

func bootLog() error {
	var ok bool
	var sysLog conf.Log
	var defLog conf.Log
	var err error
	if sysLog, ok = app.Config.Log["sys"]; !ok {
		sysLog = conf.DefaultLog
	}
	if defLog, ok = app.Config.Log["def"]; !ok {
		defLog = conf.DefaultLog
	}

	if glog2.Sys, err = newLog(sysLog); err != nil {
		return fmt.Errorf("create sys log failed: %w", err)
	}

	if glog2.Default, err = newLog(defLog); err != nil {
		return fmt.Errorf("create default log failed: %w", err)
	}

	logBootInfo("log module init")
	return nil
}

func newLog(log conf.Log) (*logrus.Logger, error) {
	logger := logrus.New()
	sysLevel, err := logrus.ParseLevel(log.Level)
	if err != nil {
		return nil, err
	}

	logger.SetLevel(sysLevel)
	if log.Format == "json" {
		logger.SetFormatter(&logrus.TextFormatter{})
	} else {
		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	switch log.Write {
	case "stderr":
		logger.SetOutput(os.Stderr)
	case "file":
		if log.FilePath == "" {
			return nil, errors.New("sys log file path no set")
		}
		f, err := os.Create(log.FilePath)
		if err != nil {
			return nil, fmt.Errorf("open log file %s failed: %w", log.FilePath, err)
		}
		logger.SetOutput(f)
	default:
		return nil, fmt.Errorf("no supported write: %s", log.Write)
	}
	return logger, nil
}

func bootDB() error {
	for k, v := range app.Config.Databases {
		gdb2.ConnConfigs[k] = gdb2.MySQLConf{
			Host:     v.Host,
			Port:     strconv.Itoa(v.Port),
			Username: v.Username,
			Password: v.Password,
			Database: v.Database,
		}
	}
	if err := gdb2.InitAll(); err != nil {
		return fmt.Errorf("init gdb module failed: %w", err)
	}

	for k, v := range app.Config.Redis {
		gredis2.Configs[k] = gredis2.Conf{
			Host:     v.Host,
			Password: v.Password,
			Port:     v.Port,
			Database: v.Index,
		}
	}

	// load migrate files
	migrate2.DB = gdb2.Def()
	if err := migrate2.InitMigrationTable(); err != nil {
		return fmt.Errorf("migrate.InitMigrationTable() failed %w", err)
	}
	return nil
}

func bootStorage() {
	storage2.Init(app.StoragePath)
	logBootInfo("storage module init")
}

func bootHTTP() {
	logBootInfo("middleware module init")
	if !app.RunningInConsole() {
		// 注册路由
		app.SetRouter(routes.NewRouter(
			routes.WithWriter(glog2.Sys.Writer()),
			routes.WithErrWriter(glog2.Sys.Writer())),
		)

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
