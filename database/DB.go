package database

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-framework/config"
	"log"
)

var Conn *gorm.DB

func Init() {
	buffer := bytes.Buffer{}
	mysqlConfig := config.App.Database.Connections.Mysql
	buffer.WriteString(mysqlConfig.Username)
	buffer.WriteString(":")
	buffer.WriteString(mysqlConfig.Password)
	buffer.WriteString("@(")
	buffer.WriteString(mysqlConfig.Host)
	buffer.WriteString(")/")
	buffer.WriteString(mysqlConfig.Database)
	buffer.WriteString("?charset=utf8&parseTime=True&loc=Local")
	conn, err := gorm.Open(config.App.Database.Default, buffer.String())
	if err != nil {
		log.Fatal("Connect mysql failed.", err)
	}
	Conn = conn
	Conn.LogMode(config.App.Debug)
}

func Close() {
	Conn.Close()
}
