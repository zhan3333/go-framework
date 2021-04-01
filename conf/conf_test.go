package conf_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/conf"
	"os"
	"testing"
)

func TestVal(t *testing.T) {
	assert.Equal(t, conf.Name, os.Getenv("APP_NAME"))
	assert.True(t, conf.Debug)
}
