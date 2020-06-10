package db_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/bootstrap"
	"go-framework/pkg/db"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap.Bootstrap()
	bootstrap.SetInTest()
	m.Run()
}

func TestConn(t *testing.T) {
	err := db.Def().DB().Ping()
	assert.Nil(t, err)
}

func TestNotExistsConn(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("must panic")
		}
	}()
	_ = db.Conn("not_exists").DB().Ping()
}
