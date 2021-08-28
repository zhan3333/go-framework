package conf

import (
	"go-framework/app"
	"go-framework/pkg/glog"
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
