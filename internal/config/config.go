package config

import (
	"time"

	"go-framework/pkg/redis"
)

type Config struct {
	App   App           `mapstructure:"app"`
	HTTP  HTTP          `mapstructure:"http"`
	JWT   JWT           `mapstructure:"jwt"`
	DB    DB            `mapstructure:"db"`
	Redis redis.Options `mapstructure:"redis"`
	Log   Log           `mapstructure:"log"`
}

type HTTP struct {
	Port    int
	Host    string
	Timeout time.Duration
}

type App struct {
	Name  string
	URL   string `json:"url"`
	Env   string
	Debug bool
}

type JWT struct {
	Secret string
	TTL    *time.Duration
	Issuer string
}

type DB struct {
	Host     string         `mapstructure:"host" json:"host"`
	Port     string         `mapstructure:"port" json:"port"`
	Username string         `mapstructure:"username" json:"username"`
	Password string         `mapstructure:"password" json:"password"`
	Database string         `mapstructure:"database" json:"database"`
	Timeout  *time.Duration `mapstructure:"timeout" json:"timeout"`
}

type Redis struct {
	Host     string `json:"host" mapstructure:"host"`
	Password string `json:"password" mapstructure:"password"`
	Port     int    `json:"port" mapstructure:"port"`
	Database int    `json:"database" mapstructure:"database"`
}

type Log struct {
	// Level: debug, info ...
	Level string `json:"level" mapstructure:"level"`
	// Format: json or text
	Format string `json:"format" mapstructure:"format"`
}
