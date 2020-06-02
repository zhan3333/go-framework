package cron

import (
	log "github.com/sirupsen/logrus"
	"go-framework/pkg/cron"
)

func Register() {
	// 每3分钟重试一次失败任务列表
	_, _ = cron.C.AddFunc("0 */1 * * * *", func() {
		log.Info("cron run")
	})
}

func Start() {
	cron.C.Start()
}

func Stop() {
	cron.C.Stop()
}
