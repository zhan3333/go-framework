package Models

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-framework/config"
	"log"
)

type db struct {
	Conn *gorm.DB
}

func (db *db) Init() {
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
	db.Conn = conn
	db.Conn.LogMode(config.App.Debug)
	db.Conn.AutoMigrate(&User{})
}

func (db *db) Close() {
	_ = db.Conn.Close()
}

var DB = db{Conn: nil}
