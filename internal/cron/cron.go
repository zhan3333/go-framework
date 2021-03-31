package cron

import (
	"go-framework/pkg/cron"
)

func Register() {
	// 每分钟运行一次
	_, _ = cron.C.AddFunc("0 */1 * * * *", func() {
		//log.Info("cron run")
	})
}

func Start() {
	cron.C.Start()
}

func Stop() {
	cron.C.Stop()
}
