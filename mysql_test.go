package main_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhan3333/gdb/v2"
	"go-framework/conf"
	"testing"
)

func TestDBPing(t *testing.T) {
	gdb.ConnConfigs = conf.Database.MySQL
	_, err := gdb.InitDef()
	assert.Nil(t, err)
	assert.Nil(t, gdb.DB.SQLDB.Ping())
}
