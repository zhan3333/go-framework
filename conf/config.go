package conf

import (
	gin2 "github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-framework/app"
	"go-framework/conf/env"
	_ "go-framework/conf/env"
	"os"
	"path/filepath"
	"strings"
)

var (
	GinModel string
	Name     = os.Getenv("APP_NAME")
	Url      = os.Getenv("APP_URL")
	Env      = os.Getenv("APP_ENV")
	Debug    = env.DefaultGetBool("DEBUG", false)
	Host     = os.Getenv("APP_HOST")
	Database = database{
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
		Redis: map[string]RedisConf{
			"default": {
				Host:     env.DefaultGet("REDIS_HOST", "127.0.0.1").(string),
				Password: env.DefaultGet("REDIS_PASSWORD", "").(string),
				Port:     env.DefaultGetInt("REDIS_PORT", 6379),
				Database: env.DefaultGetInt("REDIS_DATABASE", 0),
			},
		},
	}
	Filesystems = filesystems{
		Default: "local",
		Cloud:   "",
		Disks: Disks{
			Local: Local{
				Driver: "local",
				Root:   "app/public",
			},
		},
	}
	Logging = struct {
		Channels logs
		Default  string
	}{
		Channels: logs{
			"single": Log{
				Driver: "single",
				Path:   filepath.Join(app.StoragePath, "logs/main.log"),
				Level:  log.DebugLevel,
			},
			"gin": Log{
				Driver: "single",
				Path:   filepath.Join(app.StoragePath, "logs/route.log"),
				Level:  log.DebugLevel,
			},
			"db": Log{
				Driver: "single",
				Path:   filepath.Join(app.StoragePath, "logs/db.log"),
				Level:  log.DebugLevel,
			},
		},
		Default: "single",
	}
)

func Init() {
	if !strings.EqualFold(Env, "local") && !strings.EqualFold(Env, "production") && !strings.EqualFold(Env, "testing") {
		panic("env APP_ENV must be: local, production, testing")
	}
	switch Env {
	case "testing":
		GinModel = gin2.TestMode
	case "local":
		GinModel = gin2.DebugMode
	case "production":
		GinModel = gin2.ReleaseMode
	}
}
