package conf

import (
	gin2 "github.com/gin-gonic/gin"
	"go-framework/app"
	"go-framework/conf/env"
	_ "go-framework/conf/env"
	"os"
	"strings"
)

type Config struct {
	GinModel    string
	Name        string
	Url         string
	Env         string
	Debug       bool
	Host        string
	Database    database
	Filesystems filesystems
	Logging     logging
}

var Conf *Config

func init() {
	Conf = &Config{
		"",
		os.Getenv("APP_NAME"),
		os.Getenv("APP_URL"),
		os.Getenv("APP_ENV"),
		env.DefaultGetBool("DEBUG", false),
		os.Getenv("APP_HOST"),
		database{
			Default: os.Getenv("DB_CONNECTION"),
			Connections: connections{
				Mysql{
					Host:     os.Getenv("DB_HOST"),
					Port:     os.Getenv("DB_PORT"),
					Username: os.Getenv("DB_USERNAME"),
					Password: os.Getenv("DB_PASSWORD"),
					Database: os.Getenv("DB_DATABASE"),
				},
			},
			Redis: redis{
				RedisDefault: redisDefault{
					Host:     env.DefaultGet("REDIS_HOST", "127.0.0.1").(string),
					Password: env.DefaultGet("REDIS_PASSWORD", "").(string),
					Port:     env.DefaultGetInt("REDIS_PORT", 6379),
					Database: env.DefaultGetInt("REDIS_DATABASE", 0),
				},
			},
		},
		filesystems{
			Default: "local",
			Cloud:   "",
			Disks: Disks{
				Local: Local{
					Driver: "local",
					Root:   "app/public",
				},
			},
		},
		logging{
			Gin: gin{
				Log: ginLog{
					Path: app.StoragePath("/logs/gin.log"),
				},
			},
		},
	}
	if !strings.EqualFold(Conf.Env, "local") && !strings.EqualFold(Conf.Env, "production") && !strings.EqualFold(Conf.Env, "testing") {
		panic("env APP_ENV must be: local, production, testing")
	}
	switch Conf.Env {
	case "testing":
		Conf.GinModel = gin2.TestMode
	case "local":
		Conf.GinModel = gin2.DebugMode
	case "production":
		Conf.GinModel = gin2.ReleaseMode
	}
}
