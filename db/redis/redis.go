package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"go-framework/conf"
)

var client *redis.Client

func Init() {
	c := conf.Conf.Database.Redis.RedisDefault
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.Database,
	})
	fmt.Println(c.Password)

	_, err := client.Ping().Result()
	if err != nil {
		_ = fmt.Errorf("Redis can't connect %s:%d \n", c.Host, c.Port)
	}
}

func Client() *redis.Client {
	if client == nil {
		panic("Redis need call Init() before use.")
	}
	return client
}

func Close() {
	if client != nil {
		_ = client.Close()
	}
}
