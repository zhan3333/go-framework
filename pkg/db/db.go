package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Options struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Timeout  *time.Duration
}

func (o *Options) String() string {
	str := fmt.Sprintf("%s:%s@(%s:%s)/%s?",
		o.Username, o.Password, o.Host, o.Port, o.Database)
	if o.Timeout != nil {
		str = fmt.Sprintf("%s&timeout=%ds", str, int(o.Timeout.Seconds()))
	}
	str += "&parseTime=True&charset=utf8mb4"
	return str
}

func New(options *Options) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(options.String()), nil)
}
