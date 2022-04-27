package repository

import "context"

type ClientRepository interface {
	Subscribe(ctx context.Context, msgChan chan []byte, key string)
	Publish(ctx context.Context, key string, msg string)
}
