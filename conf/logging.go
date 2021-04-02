package conf

import (
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
			Level:  glog.DebugLevel,
		},
		"route": {
			Driver: glog.SINGLE,
			Path:   filepath.Join(app.StoragePath, "logs/route.log"),
			Level:  glog.DebugLevel,
		},
		"db": {
			Driver: glog.SINGLE,
			Path:   filepath.Join(app.StoragePath, "logs/db.log"),
			Level:  glog.DebugLevel,
		},
		"cron": {
			Driver: glog.SINGLE,
			Path:   filepath.Join(app.StoragePath, "logs/cron.log"),
			Level:  glog.DebugLevel,
		},
	},
	Default: glog.DefLogChannel,
}
