package gredis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
)

type Conf struct {
	Host     string
	Password string
	Port     int
	Database int
}

var connections = map[string]*redis.Client{}
var DefaultConn = "default"
var Configs = make(map[string]Conf)

func initConn(c Conf) (*redis.Client, error) {
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
	if c, ok := Configs[name]; ok {
		connections[name], _ = initConn(c)
		return connections[name]
	}
	panic(fmt.Sprintf("Not config redis connection: %s", name))
}

// Get default connection
func Def() *redis.Client {
	return Conn("default")
}

// reset settings and connections
func Reset() {
	for k, conn := range connections {
		if err := conn.Close(); err != nil {
			log.Printf("Close mysql conn %s err: %+v", k, err)
		}
	}
	connections = map[string]*redis.Client{}
}

func Close() {
	for k, conn := range connections {
		if err := conn.Close(); err != nil {
			log.Printf("Close mysql conn %s err: %+v", k, err)
		}
	}
}
