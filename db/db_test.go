package db_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/bootstrap"
	"go-framework/db"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap.Bootstrap()
	bootstrap.SetInTest()
	m.Run()
}

func TestConn(t *testing.T) {
	err := db.Init()
	assert.Nil(t, err)
	err = db.Conn.DB().Ping()
	assert.Nil(t, err)
}
