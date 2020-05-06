package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-framework/conf"
	"log"
)

var Conn *gorm.DB

func Init() {
	mc := conf.Conf.Database.Connections.Mysql
	connStr := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", mc.Username, mc.Password, mc.Host, mc.Database)
	conn, err := gorm.Open(conf.Conf.Database.Default, connStr)
	if err != nil {
		log.Fatalf("Connect mysql failed: %s", err.Error())
	}
	Conn = conn
	Conn.LogMode(conf.Conf.Debug)
}

func Close() {
	err := Conn.Close()
	if err != nil {
		fmt.Printf("Close db conn err: %s", err.Error())
	}
}
