package main

import (
	"go-framework/bootstrap"
	"go-framework/conf"
	"go-framework/routes"
	"log"
)

func main() {
	defer bootstrap.Destroy()
	bootstrap.Bootstrap()
	router := routes.GetRouter()
	err := router.Run(conf.Conf.Host)
	log.Println(err)
}
