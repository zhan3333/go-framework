package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"go-framework/conf"
)

var connections = map[string]*redis.Client{}

func Init() {
	var err error
	// 初始化 default 连接, 其它连接使用时才会连接
	connections["default"], err = initConn(conf.Database.Redis["default"])
	if err != nil {
		panic(fmt.Sprintf("RedisConf init connection default failed: [%+v]", err))
	}
}

func initConn(c conf.RedisConf) (*redis.Client, error) {
	var client *redis.Client
	client = newConn(fmt.Sprintf("%s:%d", c.Host, c.Port), c.Password, c.Database)
	_, err := client.Ping().Result()
	if err != nil {
		return client, err
	}
	return client, nil
}

func newConn(addr string, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// Get connection
func Conn(name string) *redis.Client {
	if client, ok := connections[name]; ok {
		return client
	}
	if c, ok := conf.Database.Redis[name]; ok {
		connections[name], _ = initConn(c)
		return connections[name]
	}
	panic(fmt.Sprintf("Not config redis connection: %s", name))
}

// Get default connection
func Client() *redis.Client {
	if client, ok := connections["default"]; ok {
		return client
	}
	panic("RedisConf need call Init() before use.")
}

func Close() {
	for _, v := range connections {
		_ = v.Close()
	}
}
