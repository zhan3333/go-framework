package redis_test

import (
	redis2 "github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	"go-framework/bootstrap"
	"go-framework/db/redis"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	bootstrap.SetInTest()
	bootstrap.Bootstrap()
	m.Run()
}

func TestPing(t *testing.T) {
	pong, err := redis.Client().Ping().Result()
	assert.Nil(t, err)
	t.Logf("Pong %s", pong)
}

func TestGetSet(t *testing.T) {
	// test exists key
	err := redis.Client().Set("test", "test", 0*time.Second).Err()
	assert.Nil(t, err)
	ret, err := redis.Client().Get("test").Result()
	assert.Equal(t, "test", ret)
	assert.Nil(t, err)

	err = redis.Client().Del("test").Err()
	assert.Nil(t, err)

	// test not exists key
	ret2, err := redis.Client().Get("test_no_exists").Result()
	assert.Equal(t, "", ret2)
	assert.NotNil(t, err)
	assert.IsType(t, redis2.Nil, err)
}
