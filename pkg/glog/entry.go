package glog

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Entry struct {
	// channel config
	Config Log
	// channel output
	Write io.Writer
	*logrus.Logger
}

func NewEntry(config Log) *Entry {
	var err error
	var format logrus.Formatter
	entry := &Entry{
		config,
		nil,
		logrus.New(),
	}
	if config.LogFormatter != nil {
		format = config.LogFormatter
	} else {
		format = DefaultFormat
	}
	entry.SetLevel(config.Level)
	if config.ReportCall {
		entry.SetReportCaller(true)
	}
	// add hooks
	for _, h := range config.Hooks {
		entry.AddHook(h)
	}

	if config.Driver == NONE {
		config.Driver = STDOUT
	}
	switch config.Driver {
	case STDOUT:
		break
	case DAILY:
		// daily driver
		dailyHook, write, err := newDailyHook(config.Path, time.Duration(config.Days)*time.Hour*24, format)
		entry.Write = write
		if err != nil {
			logrus.WithError(err).Errorf("glog channel path [%s] failed", config.Path)
		}
		entry.AddHook(dailyHook)
		entry.SetOutput(ioutil.Discard)
	case SINGLE:
		entry.SetFormatter(format)
		// create dir
		err = os.MkdirAll(filepath.Dir(config.Path), os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Create dir %s failed: [%+v]", filepath.Dir(config.Path), err))
		}
		// create log file
		f, err := os.OpenFile(config.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("Create file %s failed: [%+v]", config.Path, err))
		}
		openFiles = append(openFiles, f)
		entry.SetOutput(f)
		entry.Write = f
	}
	return entry
}

func newDailyHook(logFileName string, maxAge time.Duration, format logrus.Formatter) (*lfshook.LfsHook, io.Writer, error) {
	logExt := filepath.Ext(logFileName)
	logName := strings.TrimRight(logFileName, logExt)
	writer, err := rotatelogs.New(
		logName+"-%Y-%m-%d"+filepath.Ext(logExt),
		rotatelogs.WithLinkName(logName),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithMaxAge(maxAge),
	)

	if err != nil {
		return nil, nil, err
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		DebugLevel: writer,
		InfoLevel:  writer,
		WarnLevel:  writer,
		ErrorLevel: writer,
		FatalLevel: writer,
		PanicLevel: writer,
	}, format)
	return lfsHook, writer, nil
}
