package conf

import (
	log "github.com/sirupsen/logrus"
	"github.com/zhan3333/glog"
	"go-framework/app"
	"path/filepath"
)

var Logging = struct {
	Channels map[string]glog.Log
	Default  string
}{
	Channels: map[string]glog.Log{
		glog.DefLogChannel: {
			Driver: glog.SINGLE,
			Path:   filepath.Join(app.StoragePath, "logs/main.log"),
			Level:  log.DebugLevel,
		},
		"gin": {
			Driver: glog.SINGLE,
			Path:   filepath.Join(app.StoragePath, "logs/route.log"),
			Level:  log.DebugLevel,
		},
		"db": {
			Driver: glog.SINGLE,
			Path:   filepath.Join(app.StoragePath, "logs/db.log"),
			Level:  log.DebugLevel,
		},
		"cron": {
			Driver: glog.SINGLE,
			Path:   filepath.Join(app.StoragePath, "logs/cron.log"),
			Level:  log.DebugLevel,
		},
	},
	Default: glog.DefLogChannel,
}
