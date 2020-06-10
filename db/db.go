package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"go-framework/conf"
	"go-framework/pkg/glog"
)

var connections = map[string]*gorm.DB{}

func InitDef() (*gorm.DB, error) {
	return InitConn(conf.Database.MySQL["default"])
}

func InitConn(c conf.MysqlConf) (*gorm.DB, error) {
	var conn *gorm.DB
	var err error
	conn, err = gorm.Open("mysql", c.String())
	if err != nil {
		log.Errorf("Connect mysql failed: %+v", err)
		return conn, err
	}
	conn.LogMode(conf.Debug)
	conn.SetLogger(glog.Channel("db"))
	conn.DB().SetConnMaxLifetime(c.MaxLiftTime)
	return conn, nil
}

func Close() {
	for k, conn := range connections {
		if err := conn.Close(); err != nil {
			log.Printf("Close mysql conn %s err: %+v", k, err)
		}
	}
}

// get default conn
func Def() *gorm.DB {
	return Conn("default")
}

// get name conn
func Conn(name string) *gorm.DB {
	if conn, ok := connections[name]; ok {
		return conn
	}
	if c, ok := conf.Database.MySQL[name]; ok {
		connections[name], _ = InitConn(c)
		return connections[name]
	}
	panic(fmt.Sprintf("Can't read mysql config: %s", name))
}

// 创建一个数据库连接
func New(c conf.MysqlConf) (*gorm.DB, error) {
	conn, err := gorm.Open("mysql", c.String())
	if err != nil {
		log.Errorf("Connect mysql failed: %s", err.Error())
		return conn, err
	}
	return conn, nil
}
