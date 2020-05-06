package log

import (
	log "github.com/sirupsen/logrus"
	"go-framework/app"
	"go-framework/conf"
	"os"
)

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	//i := os.Stdout
	f, _ := os.Create(app.StoragePath("logs/main.log"))
	log.SetOutput(f)
	if conf.Conf.Debug {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
