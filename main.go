package main

import (
	"github.com/gin-gonic/gin"
	"go-websocket-demo/src/infrastructure/redis"
	"go-websocket-demo/src/infrastructure/server"
	"go-websocket-demo/src/interfaces/controller"
	"go-websocket-demo/src/resistory"
)

func main() {
	redis := redis.New()

	repositoryRegistry := resistory.NewRepository()
	usecaseRegistry := resistory.NewUsecase()

	ctrl := controller.NewController(redis, repositoryRegistry, usecaseRegistry)
	var middlewares []gin.HandlerFunc
	router := server.NewRouter(middlewares, ctrl)
	server.Run(router)
}
