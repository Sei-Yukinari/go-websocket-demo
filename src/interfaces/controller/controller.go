package controller

import (
	"go-websocket-demo/src/infrastructure/redis"
	"go-websocket-demo/src/resistory"
)

type Controller struct {
	redis      *redis.Client
	repository resistory.Repository
	usecase    resistory.Usecase
}

func NewController(redis *redis.Client,
	repository resistory.Repository,
	usecase resistory.Usecase,
) *Controller {
	return &Controller{
		redis:      redis,
		repository: repository,
		usecase:    usecase,
	}
}
