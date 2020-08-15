package conf

import (
	"github.com/zhan3333/gdb"
	"go-framework/conf/env"
	"os"
	"time"
)

type RedisConf struct {
	Host     string
	Password string
	Port     int
	Database int
}

type database struct {
	MySQL map[string]gdb.MysqlConf
	Redis map[string]RedisConf
}

var Database = database{
	MySQL: map[string]gdb.MysqlConf{
		"default": {
			Host:        os.Getenv("DB_HOST"),
			Port:        os.Getenv("DB_PORT"),
			Username:    os.Getenv("DB_USERNAME"),
			Password:    os.Getenv("DB_PASSWORD"),
			Database:    os.Getenv("DB_DATABASE"),
			MaxLiftTime: time.Second * 60,
			LogMode:     env.DefaultGetBool("DB_LOG_MODE", true),
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
