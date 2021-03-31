package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"runtime"
)

// 储存全局变量

var (
	InTest      bool
	InConsole   bool
	Path        string
	TestPath    string
	StoragePath string
	// 是否引导完毕
	Booted bool
	router *gin.Engine
)

func init() {
	// 初始化项目的路径
	Path = GetBasePath()
	TestPath = fmt.Sprintf("%s/tests", Path)
	StoragePath = fmt.Sprintf("%s/storage", Path)
}

// 获取项目基础路径的绝对值
func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}

func RunningInConsole() bool {
	return InConsole
}

func RunningInTest() bool {
	return InTest
}

func IsBooted() bool {
	return Booted
}

func SetRouter(engine *gin.Engine) {
	router = engine
}

func GetRouter() *gin.Engine {
	if router == nil {
		panic("router not set")
	}
	return router
}
