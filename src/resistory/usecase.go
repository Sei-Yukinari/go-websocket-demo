package resistory

import (
	"go-websocket-demo/src/domain/repository"
	"go-websocket-demo/src/usecase/client"
)

type Usecase interface {
	NewClient(repository.ClientRepository) client.Connect
}

type usecaseImpl struct{}

func NewUsecase() Usecase {
	return &usecaseImpl{}
}

func (u usecaseImpl) NewClient(r repository.ClientRepository) client.Connect {
	return client.NewConnect(r)
}
