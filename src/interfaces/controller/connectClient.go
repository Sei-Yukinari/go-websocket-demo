package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-websocket-demo/src/infrastructure/http"
	"go-websocket-demo/src/usecase/client"
)

type ConnectClientController struct {
	controller *Controller
}

func NewConnectClientController(controller *Controller) *ConnectClientController {
	return &ConnectClientController{
		controller: controller,
	}
}

func (connClient *ConnectClientController) Connect(c *gin.Context) {
	ctx := c.Request.Context()
	repo := connClient.controller.repository.NewClient(connClient.controller.redis)

	w := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := w.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		c.JSON(500, http.ResponseFormat{
			StatusCode: 500,
			Message:    "Failed to create connection",
		})
	}
	username := c.Query("name")
	chatbox := c.Query("chatbox")
	if username == "" || chatbox == "" {
		c.JSON(400, http.ResponseFormat{
			StatusCode: 400,
			Message:    "query name or chatbox must be filled",
		})
	}
	in := &client.ConnInput{
		Conn: conn, Username: username, Chatbox: chatbox,
	}
	connClient.controller.usecase.
		NewClient(repo).
		Invoke(ctx, in)
}
