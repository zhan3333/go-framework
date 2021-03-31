package conf_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/conf"
	"go-framework/core/boot"
	"os"
	"testing"
)

func TestVal(t *testing.T) {
	boot.SetInTest()
	assert.Equal(t, conf.Name, os.Getenv("APP_NAME"))
	assert.True(t, conf.Debug)
}
