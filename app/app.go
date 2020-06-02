package app

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var (
	InTest      bool
	InConsole   bool
	Path        string
	TestPath    string
	StoragePath string
	IsBootstrap bool
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

func Booted() bool {
	return IsBootstrap
}
