package app_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/app"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestGetBasePath(t *testing.T) {
	assert.NotEmpty(t, app.GetBasePath())
}

func TestRunningInConsole(t *testing.T) {
	assert.False(t, app.RunningInConsole())
	app.InConsole = true
	assert.True(t, app.RunningInConsole())
}

func TestRunningInTest(t *testing.T) {
	assert.False(t, app.RunningInTest())
	app.InTest = true
	assert.True(t, app.RunningInTest())
}

func TestBooted(t *testing.T) {
	assert.False(t, app.Booted())
	app.IsBootstrap = true
	assert.True(t, app.Booted())
}
