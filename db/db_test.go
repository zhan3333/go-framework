package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConn(t *testing.T) {
	Init()
	err := Conn.DB().Ping()
	assert.Nil(t, err)
}
