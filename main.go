package main

import (
	"go-framework/app"
	"go-framework/conf"
	"go-framework/core/boot"
	// 引入的形式启动框架
	_ "go-framework/core/boot/http"
	"log"
)

// @title go-framework
// @version 1.0
// @description gin framework
// @license.name none
func main() {
	defer boot.Destroy()
	err := app.GetRouter().Run(conf.Host)
	log.Println(err)
}
