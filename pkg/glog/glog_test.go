package glog_test

import (
	log "github.com/sirupsen/logrus"
	"go-framework/boot"
	"testing"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	m.Run()
}

func TestLog(t *testing.T) {
	log.WithFields(log.Fields{
		"aniaml": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.Print("test \n")
}
