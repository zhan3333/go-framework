package main

import (
	"go-framework/boot"
	"go-framework/conf"
	routes "go-framework/internal/route"
	"log"
)

// @title go-framework
// @version 1.0
// @description gin framework
// @license.name none
func main() {
	defer boot.Destroy()
	boot.Boot()
	router := routes.GetRouter()
	err := router.Run(conf.Host)
	log.Println(err)
}
