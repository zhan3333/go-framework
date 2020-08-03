package conf

import (
	"github.com/zhan3333/gdb"
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
