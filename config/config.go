package config

import (
	"os"
)

type Mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type connections struct {
	Mysql
}

type redis struct {
	RedisDefault redisDefault
}

type redisDefault struct {
	Host     string
	Password string
	Port     string
	Database int
}

type database struct {
	Default     string
	Connections connections
	Redis       redis
}

type Config struct {
	Name     string
	Url      string
	Env      string
	Debug    bool
	Host     string
	Database database
}

var App *Config

func Init() {
	App = &Config{
		os.Getenv("APP_NAME"),
		os.Getenv("APP_URL"),
		os.Getenv("APP_ENV"),
		DefaultGetBool("DEBUG", false),
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
					Host:     DefaultGet("REDIS_HOST", "127.0.0.1").(string),
					Password: DefaultGet("REDIS_PASSWORD", nil).(string),
					Port:     DefaultGet("REDIS_PORT", 6379).(string),
					Database: DefaultGetInt("REDIS_DATABASE", 0),
				},
			},
		},
	}
}
