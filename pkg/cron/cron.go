package cron

import (
	cron2 "github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"go-framework/pkg/glog"
)

var C *cron2.Cron

type Logger struct {
}

func (Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	glog.Channel("cron").WithFields(log.Fields{
		"err":           err,
		"keysAndValues": keysAndValues,
	}).Error(msg)
}

func (Logger) Info(msg string, keysAndValues ...interface{}) {
	glog.Channel("cron").WithFields(log.Fields{
		"keysAndValues": keysAndValues,
	}).Info(msg)
}

func init() {
	C = cron2.New(
		cron2.WithSeconds(),
		cron2.WithLogger(Logger{}),
		cron2.WithChain(cron2.SkipIfStillRunning(Logger{})),
	)
}
