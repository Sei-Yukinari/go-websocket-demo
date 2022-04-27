package resistory

import (
	"go-websocket-demo/src/domain/repository"
	"go-websocket-demo/src/gateway"
	"go-websocket-demo/src/infrastructure/redis"
)

type repositoryImpl struct{}

type Repository interface {
	NewClient(redis *redis.Client) repository.ClientRepository
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) NewClient(redis *redis.Client) repository.ClientRepository {
	return gateway.NewClient(redis)
}
