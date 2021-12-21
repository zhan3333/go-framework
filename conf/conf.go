package conf

import (
	"time"
)

type Config struct {
	App       App
	HTTP      HTTP
	JWT       JWT
	Databases map[string]DB    `mapstructure:"db"`
	Redis     map[string]Redis `mapstructure:"redis"`
	Log       map[string]Log   `mapstructure:"log"`
	Cron      Cron
}

type HTTP struct {
	Port         int
	Host         string
	Timeout      time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

type Cron struct {
	Enable bool
}

type App struct {
	Name  string
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
	// text / json (默认)
	Format string
}

var DefaultLog = Log{
	Write:  "stderr",
	Level:  "info",
	Format: "json",
}
