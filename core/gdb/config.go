package gdb

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type MySQLConf struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Timeout  *time.Duration
	// 是否解析时间, 默认值为 True
	ParseTime    *bool
	Charset      *string
	Loc          *string
	GORMConfig   *gorm.Config
	MaxLiftTime  *time.Duration
	MaxOpenConns *int
	MaxIdleConns *int
}

type Logger interface {
	Print(v ...interface{})
}

func (c MySQLConf) String() string {
	str := fmt.Sprintf("%s:%s@(%s:%s)/%s?",
		c.Username, c.Password, c.Host, c.Port, c.Database)
	if c.Timeout != nil {
		str = fmt.Sprintf("%s&timeout=%ds", str, int(c.Timeout.Seconds()))
	}
	bParseTime := "True"
	if c.ParseTime != nil && !*c.ParseTime {
		bParseTime = "False"
	}
	str = fmt.Sprintf("%s&parseTime=%s", str, bParseTime)
	if c.Charset != nil {
		str = fmt.Sprintf("%s&charset=%s", str, *c.Charset)
	}
	if c.Loc != nil {
		str = fmt.Sprintf("%s&loc=%s", str, *c.Loc)
	}
	return str
}
