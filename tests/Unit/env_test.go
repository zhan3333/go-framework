package Unit

import (
	"github.com/stretchr/testify/assert"
	"go-framework/bootstrap"
	"go-framework/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap.SetInTest()
	bootstrap.Bootstrap()
	m.Run()
}

func TestDefaultGet(t *testing.T) {
	t.Log(config.App.Name)
	assert.Equal(t, config.App.Name, os.Getenv("APP_NAME"))
}
