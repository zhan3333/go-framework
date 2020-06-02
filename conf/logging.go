package conf

import log "github.com/sirupsen/logrus"

type logs map[string]Log

type Log struct {
	Driver string
	Path   string
	Level  log.Level
	Days   int
}
