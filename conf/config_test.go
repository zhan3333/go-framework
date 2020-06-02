package conf_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/bootstrap"
	"go-framework/conf"
	"os"
	"testing"
)

func TestVal(t *testing.T) {
	bootstrap.SetInTest()
	assert.Equal(t, conf.Name, os.Getenv("APP_NAME"))
}
