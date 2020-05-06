package app

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type application struct {
	InTest      bool
	AppPath     string
	TestPath    string
	StoragePath string
}

var App = application{
	InTest:      false,
	AppPath:     "",
	TestPath:    "",
	StoragePath: "",
}

func Init() {
	appPath := GetBasePath()
	App.AppPath = appPath
	App.AppPath = fmt.Sprintf("%s/tests", appPath)
	App.StoragePath = fmt.Sprintf("%s/storage", appPath)
}

func StoragePath(path string) string {
	return filepath.Join(App.StoragePath, path)
}

//func TestPath(path string) string {
//	return filepath.Join(App.TestPath, path)
//}
//
//func AppPath(path string) string {
//	return filepath.Join(App.AppPath, path)
//}

// 获取项目基础路径的绝对值
func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}
