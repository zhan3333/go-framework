package logger

import (
	"github.com/sirupsen/logrus"
	"io"
)

type Options struct {
	Out    io.Writer
	Format logrus.Formatter
	Level  logrus.Level
}

func New(options *Options) *logrus.Logger {
	var log = &logrus.Logger{
		Out:       options.Out,
		Formatter: options.Format,
		Hooks:     make(logrus.LevelHooks),
		Level:     options.Level,
	}
	return log
}
