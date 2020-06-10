package redis_test

import (
	redis2 "github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	"go-framework/boot"
	"go-framework/conf"
	"go-framework/pkg/redis"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	m.Run()
}

func TestPing(t *testing.T) {
	pong, err := redis.Def().Ping().Result()
	assert.Nil(t, err)
	t.Logf("Pong %s", pong)
}

func TestGetSet(t *testing.T) {
	// test exists key
	err := redis.Def().Set("test", "test", 0*time.Second).Err()
	assert.Nil(t, err)
	ret, err := redis.Def().Get("test").Result()
	assert.Equal(t, "test", ret)
	assert.Nil(t, err)

	err = redis.Def().Del("test").Err()
	assert.Nil(t, err)

	// test not exists key
	ret2, err := redis.Def().Get("test_no_exists").Result()
	assert.Equal(t, "", ret2)
	assert.NotNil(t, err)
	assert.IsType(t, redis2.Nil, err)
}

func TestNewConn(t *testing.T) {
	conf.Database.Redis["new"] = conf.RedisConf{
		Host:     "127.0.0.1",
		Password: "",
		Port:     6379,
		Database: 0,
	}
	pong, err := redis.Conn("new").Ping().Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", pong)
}

func TestNotExistsConn(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	redis.Conn("not_exists")
}
