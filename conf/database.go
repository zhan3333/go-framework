package conf

import (
	"fmt"
	"time"
)

type MysqlConf struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Database    string
	MaxLiftTime time.Duration
}

type RedisConf struct {
	Host     string
	Password string
	Port     int
	Database int
}

func (c MysqlConf) String() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=15s",
		c.Username, c.Password, c.Host, c.Port, c.Database)
}

type database struct {
	MySQL map[string]MysqlConf
	Redis map[string]RedisConf
}
