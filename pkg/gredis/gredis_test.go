package gredis_test

import (
	redis2 "github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	"go-framework/pkg/gredis"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	gredis.Configs = map[string]gredis.Conf{
		gredis.DefaultConn: {
			Host:     "127.0.0.1",
			Password: "",
			Port:     6379,
			Database: 0,
		},
	}
	m.Run()
}

func TestPing(t *testing.T) {
	pong, err := gredis.Def().Ping().Result()
	assert.Nil(t, err)
	t.Logf("Pong %s", pong)
	gredis.Close()
}

func TestGetSet(t *testing.T) {
	// test exists key
	err := gredis.Def().Set("test", "test", 0*time.Second).Err()
	assert.Nil(t, err)
	ret, err := gredis.Def().Get("test").Result()
	assert.Equal(t, "test", ret)
	assert.Nil(t, err)

	err = gredis.Def().Del("test").Err()
	assert.Nil(t, err)

	// test not exists key
	ret2, err := gredis.Def().Get("test_no_exists").Result()
	assert.Equal(t, "", ret2)
	assert.NotNil(t, err)
	assert.IsType(t, redis2.Nil, err)
}

func TestNewConn(t *testing.T) {
	gredis.Configs = map[string]gredis.Conf{
		gredis.DefaultConn: {
			Host:     "127.0.0.1",
			Password: "",
			Port:     6379,
			Database: 0,
		},
		"new": {
			Host:     "127.0.0.1",
			Password: "",
			Port:     6379,
			Database: 0,
		},
	}
	gredis.Reset()
	pong, err := gredis.Conn("new").Ping().Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", pong)
}

func TestNotExistsConn(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	gredis.Conn("not_exists")
}
