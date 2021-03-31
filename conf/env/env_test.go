package env_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/conf/env"
	"go-framework/core/boot"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	m.Run()
}

func TestDefaultGet(t *testing.T) {
	assert.Equal(t, os.Getenv("APP_NAME"), env.DefaultGet("APP_NAME", ""))
}
