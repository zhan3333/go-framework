package glog

import (
	"github.com/sirupsen/logrus"
	"os"
)

// logrus 的包装工具
// DefaultFormat 配置默认的日志格式
// LogConfigs 日志通道配置

const (
	PanicLevel logrus.Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

const (
	NONE = iota
	// SINGLE 单文件驱动
	SINGLE
	// DAILY 日驱动
	DAILY
	STDOUT
)

var AllLevels = logrus.AllLevels

type Log struct {
	Driver       uint
	Path         string
	Level        logrus.Level
	Days         int
	LogFormatter logrus.Formatter
	ReportCall   bool
	Hooks        []logrus.Hook
}

// DefaultFormat 默认日志格式
var DefaultFormat logrus.Formatter = LocalFormatter{&logrus.JSONFormatter{
	PrettyPrint:       false,
	DisableHTMLEscape: true,
}}
var DefLogChannel = "default"
var LogConfigs = map[string]Log{
	DefLogChannel: {
		Driver:       STDOUT,
		Level:        TraceLevel,
		LogFormatter: DefaultFormat,
		ReportCall:   true,
	},
}

// 缓存渠道信息
var channels = map[string]*Entry{}

// 被打开的文件对象
var openFiles []*os.File

type LocalFormatter struct {
	logrus.Formatter
}

func (u LocalFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.Local()
	return u.Formatter.Format(e)
}

// 获取默认通道
func Def() *Entry {
	return Channel(DefLogChannel)
}

// 获取指定通道
func Channel(name string) *Entry {
	if l, ok := channels[name]; ok {
		return l
	}
	if c, ok := LogConfigs[name]; ok {
		channels[name] = NewEntry(c)
		return channels[name]
	}
	return channels[DefLogChannel]
}

// 加载所有通道
func LoadChannels() {
	// channels
	for name, logConf := range LogConfigs {
		channels[name] = NewEntry(logConf)
	}
}

// 重载所有通道, 通常用于修改了默认通道或者配置后调用
func ReloadChannels() {
	channels = map[string]*Entry{}
	for _, f := range openFiles {
		_ = f.Close()
	}
	LoadChannels()
}

// 关闭文件
func Close() {
	for _, f := range openFiles {
		_ = f.Close()
	}
}
