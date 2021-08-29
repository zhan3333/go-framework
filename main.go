package main

import (
	"fmt"
	"go-framework/app"
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
	err := app.GetRouter().Run(fmt.Sprintf("%s:%d", app.Config.App.Host, app.Config.App.Port))
	log.Println(err)
}
