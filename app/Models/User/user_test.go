package User_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/app/Models/User"
	"go-framework/bootstrap"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap.SetInTest()
	bootstrap.Bootstrap()
	m.Run()
}

func TestMailIsExists(t *testing.T) {
	exists := User.EmailIsExists("390961827@qq.com")
	assert.Equal(t, false, exists)
}
