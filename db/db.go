package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-framework/conf"
	"log"
	"time"
)

var Conn *gorm.DB

func Init() {
	mc := conf.Conf.Database.Connections.Mysql

	Conn = New(mc.Username, mc.Password, mc.Host, mc.Port, mc.Database)

	Conn.LogMode(conf.Conf.Debug)

	// 重新连接间隔
	Conn.DB().SetConnMaxLifetime(60 * time.Second)
}

func Close() {
	err := Conn.Close()
	if err != nil {
		fmt.Printf("Close db conn err: %s", err.Error())
	}
}

// 创建一个数据库连接
func New(name string, password string, host string, port string, database string) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		name, password, host, port, database)
	conn, err := gorm.Open(conf.Conf.Database.Default, connStr)
	if err != nil {
		log.Fatalf("Connect mysql failed: %s", err.Error())
	}
	return conn
}
