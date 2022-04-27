package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-websocket-demo/src/config"
)

type Client = redis.Client

func New() *Client {
	c := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.URL,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.DB,
	})

	ctx := context.Background()
	err := c.Ping(ctx).Err()
	if err != nil {
		fmt.Errorf("failed to connect redis:%v\n", err)
	}
	return c
}
