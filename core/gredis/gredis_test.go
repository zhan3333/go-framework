package gredis_test

import (
	redis2 "github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	gredis2 "go-framework/core/gredis"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	gredis2.Configs = map[string]gredis2.Conf{
		gredis2.DefaultConn: {
			Host:     "127.0.0.1",
			Password: "",
			Port:     6379,
			Database: 0,
		},
	}
	m.Run()
}

func TestPing(t *testing.T) {
	pong, err := gredis2.Def().Ping().Result()
	assert.Nil(t, err)
	t.Logf("Pong %s", pong)
	gredis2.Close()
}

func TestGetSet(t *testing.T) {
	// test exists key
	err := gredis2.Def().Set("test", "test", 0*time.Second).Err()
	assert.Nil(t, err)
	ret, err := gredis2.Def().Get("test").Result()
	assert.Equal(t, "test", ret)
	assert.Nil(t, err)

	err = gredis2.Def().Del("test").Err()
	assert.Nil(t, err)

	// test not exists key
	ret2, err := gredis2.Def().Get("test_no_exists").Result()
	assert.Equal(t, "", ret2)
	assert.NotNil(t, err)
	assert.IsType(t, redis2.Nil, err)
}

func TestNewConn(t *testing.T) {
	gredis2.Configs = map[string]gredis2.Conf{
		gredis2.DefaultConn: {
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
	gredis2.Reset()
	pong, err := gredis2.Conn("new").Ping().Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", pong)
}

func TestNotExistsConn(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	gredis2.Conn("not_exists")
}
