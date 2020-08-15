package conf_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/boot"
	"go-framework/conf"
	"os"
	"testing"
)

func TestVal(t *testing.T) {
	boot.SetInTest()
	assert.Equal(t, conf.Name, os.Getenv("APP_NAME"))
	assert.True(t, conf.Debug)
}
