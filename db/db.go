package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"go-framework/conf"
	"go-framework/pkg/glog"
	"time"
)

var Conn *gorm.DB

type dbConf struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (c dbConf) String() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=15s",
		c.Username, c.Password, c.Host, c.Port, c.Database)
}

func Init() error {
	mc := conf.Database.Connections.Mysql
	c := dbConf{
		Username: mc.Username,
		Password: mc.Password,
		Host:     mc.Host,
		Port:     mc.Port,
		Database: mc.Database,
	}
	Conn, _ = New(c)

	Conn.LogMode(conf.Debug)
	Conn.SetLogger(glog.Channel("db"))
	if err := Conn.DB().Ping(); err != nil {
		return fmt.Errorf("连接到数据库 %s 失败: [%+v]", c.String(), err)
	}
	// 重新连接间隔
	Conn.DB().SetConnMaxLifetime(60 * time.Second)
	return nil
}

func Close() {
	err := Conn.Close()
	if err != nil {
		fmt.Printf("Close db conn err: %s", err.Error())
	}
}

// 创建一个数据库连接
func New(c dbConf) (*gorm.DB, error) {
	conn, err := gorm.Open(conf.Database.Default, c.String())
	if err != nil {
		log.Errorf("Connect mysql failed: %s", err.Error())
		return conn, err
	}
	return conn, nil
}
