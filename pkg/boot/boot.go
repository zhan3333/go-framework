package boot

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"

	"go-framework/internal/auth"
	"go-framework/internal/config"
	routes "go-framework/internal/route"
	"go-framework/internal/validator"
	"go-framework/pkg/db"
	"go-framework/pkg/lgo"
	"go-framework/pkg/logger"
	"go-framework/pkg/redis"
)

type Booted struct {
	*lgo.Dependencies
	Booted bool
}

func (b *Booted) Destroy() {
	if b == nil {
		return
	}
	if b.Redis != nil {
		_ = b.Redis.Close()
	}
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

// Boot 应用启动入口
func Boot(opts ...Option) (*Booted, error) {
	var booted = &Booted{
		Dependencies: &lgo.Dependencies{},
	}
	var err error

	options := Options{
		enableRoutePrint: true,
	}
	for _, opt := range opts {
		opt(&options)
	}

	if booted.Config, err = loadConfig(options.configFile); err != nil {
		return nil, fmt.Errorf("boot config: %w", err)
	}

	level, err := logrus.ParseLevel(booted.Config.Log.Level)
	if err != nil {
		return nil, fmt.Errorf("boot log: level parse: %w", err)
	}

	var format logrus.Formatter = new(logrus.JSONFormatter)
	if booted.Config.Log.Format == "text" {
		format = new(logrus.TextFormatter)
	}

	booted.Logger = logger.New(&logger.Options{
		Out:    os.Stderr,
		Format: format,
		Level:  level,
	})

	booted.JWT = auth.NewJWT(&auth.Options{
		Secret: booted.Config.JWT.Secret,
		TTL:    booted.Config.JWT.TTL,
		Issuer: booted.Config.JWT.Issuer,
	})

	booted.DB, err = db.New(&db.Options{
		Host:     booted.Config.DB.Host,
		Port:     booted.Config.DB.Port,
		Username: booted.Config.DB.Username,
		Password: booted.Config.DB.Password,
		Database: booted.Config.DB.Database,
		Timeout:  booted.Config.DB.Timeout,
	})
	if err != nil {
		return nil, fmt.Errorf("boot db: %w", err)
	}

	booted.Redis, err = redis.New(&redis.Options{
		Host:     booted.Config.Redis.Host,
		Password: booted.Config.Redis.Password,
		Port:     booted.Config.Redis.Port,
		Database: booted.Config.Redis.Database,
	})

	if err != nil {
		return nil, fmt.Errorf("boot redis: %w", err)
	}

	booted.Server = routes.NewRouter(booted.Dependencies,
		routes.WithWriter(booted.Logger.Writer()),
		routes.WithErrWriter(booted.Logger.Writer()),
	)

	validator.Init()

	return booted, nil
}

func loadConfig(configFile string) (*config.Config, error) {
	var conf = &config.Config{}

	viper.SetDefault("log.level", logrus.InfoLevel)
	viper.SetDefault("log.format", "json")

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}
	return conf, nil
}
