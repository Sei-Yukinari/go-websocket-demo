package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go-websocket-demo/src/domain/model"
	"go-websocket-demo/src/domain/repository"
	"go-websocket-demo/src/infrastructure/http"
	"log"
)

type Connect interface {
	Invoke(context.Context, *ConnInput)
}

type connect struct {
	repository repository.ClientRepository
}

func NewConnect(repo repository.ClientRepository) Connect {
	return &connect{repo}
}

type ConnInput struct {
	Conn     *websocket.Conn
	Username string
	Chatbox  string
}

func (conn *connect) Invoke(ctx context.Context, in *ConnInput) {
	c := model.NewClient(in.Conn, in.Username, in.Chatbox)
	go conn.repository.Subscribe(ctx, c.MsgChan, c.Chatbox)
	go conn.receiveMessage(ctx, c)
	conn.sendMessage(c)
}

func (conn *connect) receiveMessage(ctx context.Context, c *model.ClientInfo) error {
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return err
		}
		mf := http.MessageFormat{
			Username: c.ClientName,
			Message:  string(msg),
		}
		if msg, err = json.Marshal(mf); err != nil {
			panic(err)
		}
		if err != nil {
			return err
		}
		fmt.Printf("publish:%v:%v\n", c.ClientName, string(msg))
		conn.repository.Publish(ctx, c.Chatbox, string(msg))
	}
}

func (conn *connect) sendMessage(c *model.ClientInfo) {
	for {
		select {
		case msg := <-c.MsgChan:
			c.Conn.WriteJSON(string(msg))
		}
	}
}
