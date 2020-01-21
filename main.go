package main

import (
	"go-framework/bootstrap"
	"go-framework/config"
	"log"
)

func main() {
	defer bootstrap.Destroy()
	router := bootstrap.Bootstrap()
	err := router.Run(config.App.Host)
	log.Println(err)
}
