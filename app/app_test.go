package app_test

import (
	"go-framework/app"
	"go-framework/bootstrap"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap.SetInTest()
	bootstrap.Bootstrap()
	m.Run()
}

func TestPath(t *testing.T) {
	t.Logf("Storage Path: %s", app.StoragePath(""))
}
