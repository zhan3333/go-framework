package conf

import (
	"time"

	// 加载 env 文件
	_ "go-framework/core/env"
)

type Config struct {
	App       App
	JWT       JWT
	Databases map[string]DB    `toml:"db"`
	Redis     map[string]Redis `toml:"redis"`
	Log       map[string]Log   `toml:"log"`
	Cron      Cron
}

type Cron struct {
	Enable bool
}

type App struct {
	Name  string
	Port  int
	Host  string
	URL   string `json:"url"`
	Env   string
	Debug bool
}

type JWT struct {
	Secret string
	TTL    time.Duration
	Issuer string
}

type DB struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	LogEnable bool
}

type Redis struct {
	Host     string
	Port     int
	Password string
	Index    int
}

type Log struct {
	// Write 可选 stderr (默认) | file
	Write    string
	Level    string
	FilePath string
}

var DefaultLog = Log{
	Write: "stderr",
	Level: "info",
}
