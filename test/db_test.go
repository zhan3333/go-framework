package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhan3333/gdb"
	"go-framework/boot"
	"testing"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	m.Run()
}

func TestPing(t *testing.T) {
	assert.Nil(t, gdb.Def().DB().Ping())
}
