package main_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/conf"
	"go-framework/pkg/gdb"
	"testing"
)

func TestDBPing(t *testing.T) {
	gdb.ConnConfigs = conf.Database.MySQL
	_, err := gdb.InitDef()
	assert.Nil(t, err)
	assert.Nil(t, gdb.DB.SQLDB.Ping())
}
