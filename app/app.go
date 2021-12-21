package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redis2 "github.com/go-redis/redis/v7"
	"path/filepath"
	"runtime"

	"go-framework/conf"
)

// 储存全局变量

var (
	InTest      bool
	InConsole   bool
	Path        string
	TestPath    string
	StoragePath string
	// Booted 是否引导完毕
	Booted bool
	router *gin.Engine
	Config *conf.Config
	redis  *redis2.Client
)

func init() {
	// 初始化项目的路径
	Path = GetBasePath()
	TestPath = fmt.Sprintf("%s/tests", Path)
	StoragePath = fmt.Sprintf("%s/storage", Path)
}

// GetBasePath 获取项目基础路径的绝对值
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

func SetRedis(r *redis2.Client) {
	redis = r
}

func Redis() *redis2.Client {
	if redis == nil {
		panic("app.redis not set")
	}
	return redis
}
