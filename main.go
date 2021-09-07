package main

import (
	"go-framework/app"
	"go-framework/core/boot"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

var defaultConfigFile = "config/local.toml"

// @title go-framework
// @version 1.0
// @description gin framework
// @license.name none
func main() {
	configFile := *kingpin.Flag("config", "load config file").Default(defaultConfigFile).String()
	kingpin.Parse()
	if err := boot.New(boot.WithConfigFile(configFile)); err != nil {
		log.Panicf("boot failed: %+v", err)
	}
	defer boot.Destroy()

	if err := app.Run(); err != nil {
		log.Panicf("server run failed: %+v", err)
	}
}
