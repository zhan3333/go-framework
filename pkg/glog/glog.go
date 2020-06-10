package glog

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-framework/conf"
	"os"
)

var Channels = map[string]*log.Logger{}
var openFiles []*os.File

type LocalFormatter struct {
	log.Formatter
}

func (u LocalFormatter) Format(e *log.Entry) ([]byte, error) {
	e.Time = e.Time.Local()
	return u.Formatter.Format(e)
}

func Default() *log.Logger {
	return Channel("default")
}

func Channel(name string) *log.Logger {
	if l, ok := Channels[name]; ok {
		return l
	}
	return Channels["default"]
}

// 加载所有通道
func loadChannels() {
	// default
	Channels["default"] = configDefaultLog()
	// channels
	for name, logConf := range conf.Logging.Channels {
		Channels[name] = configLog(logConf)
	}
}

func configDefaultLog() *log.Logger {
	if logC, ok := conf.Logging.Channels[conf.Logging.Default]; ok {
		l := log.StandardLogger()
		log.SetFormatter(LocalFormatter{&log.JSONFormatter{
			DisableHTMLEscape: true,
		}})
		f, err := os.Create(logC.Path)
		if err != nil {
			panic(fmt.Sprintf("Create file failed: [%+v]", err))
		}
		openFiles = append(openFiles, f)
		l.SetOutput(f)
		l.SetLevel(logC.Level)
		return l
	}
	return nil
}

func configLog(logConf conf.Log) *log.Logger {
	l := log.New()
	log.SetFormatter(LocalFormatter{&log.JSONFormatter{
		DisableHTMLEscape: true,
	}})
	f, err := os.Create(logConf.Path)
	if err != nil {
		panic(fmt.Sprintf("Create file failed: [%+v]", err))
	}
	openFiles = append(openFiles, f)
	l.SetOutput(f)
	l.SetLevel(logConf.Level)
	return l
}

// 全局使用默认通道

func Init() {
	loadChannels()
}

func Close() {
	for _, f := range openFiles {
		_ = f.Close()
	}
}
