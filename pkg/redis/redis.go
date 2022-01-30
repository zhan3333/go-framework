package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Options struct {
	Host     string
	Password string
	Port     int
	Database int
}

func New(o *Options) (*redis.Client, error) {
	var client *redis.Client
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", o.Host, o.Port),
		Password: o.Password,
		DB:       o.Database,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
