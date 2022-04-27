package gateway

import (
	"context"
	"fmt"
	"go-websocket-demo/src/domain/repository"
	"go-websocket-demo/src/infrastructure/redis"
)

type Client struct {
	redis *redis.Client
}

func NewClient(redis *redis.Client) *Client {
	return &Client{redis}
}

var _ repository.ClientRepository = (*Client)(nil)

func (c Client) Subscribe(ctx context.Context, msgChan chan []byte, key string) {
	subscriber := c.redis.Subscribe(ctx, key)
	fmt.Printf("subscribe-start:%s\n", subscriber.String())
	defer subscriber.Close()
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			fmt.Errorf("subscribe message error:%v", err)
		}
		msgChan <- []byte(msg.Payload)
	}
}

func (c Client) Publish(ctx context.Context, key string, msg string) {
	c.redis.Publish(ctx, key, msg)
}
