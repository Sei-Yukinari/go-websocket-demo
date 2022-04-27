package server

import (
	"github.com/gin-gonic/gin"
	"go-websocket-demo/src/interfaces/controller"
	netHttp "net/http"
)

func NewRouter(
	middlewares []gin.HandlerFunc,
	ctrl *controller.Controller,
) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	for _, m := range middlewares {
		r.Use(m)
	}

	connectClientRouter(r, ctrl)

	r.GET("", func(c *gin.Context) {
		netHttp.ServeFile(c.Writer, c.Request, "public/index.html")
	})
	return r
}

func connectClientRouter(r *gin.Engine, ctrl *controller.Controller) {
	controller := controller.NewConnectClientController(ctrl)
	r.GET("/ws", func(c *gin.Context) {
		controller.Connect(c)
	})
}
