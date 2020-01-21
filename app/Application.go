package app

import "path/filepath"

type application struct {
	InTest      bool
	AppPath     string
	TestPath    string
	StoragePath string
}

var Application = application{
	InTest:      false,
	AppPath:     "",
	TestPath:    "",
	StoragePath: "",
}

func StoragePath(path string) string {
	return filepath.Join(Application.StoragePath, path)
}

func TestPath(path string) string {
	return filepath.Join(Application.TestPath, path)
}

func AppPath(path string) string {
	return filepath.Join(Application.AppPath, path)
}
