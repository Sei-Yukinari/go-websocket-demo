package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go-websocket-demo/src/domain/repository"
	"go-websocket-demo/src/infrastructure/http"
	"log"
)

type ClientInfo struct {
	Conn       *websocket.Conn
	ClientName string
	Chatbox    string
	MsgChan    chan []byte
}

func NewClient(conn *websocket.Conn, name string, chatbox string) *ClientInfo {
	return &ClientInfo{
		Conn:       conn,
		ClientName: name,
		Chatbox:    chatbox,
		MsgChan:    make(chan []byte),
	}
}

func (c ClientInfo) Run(repo repository.ClientRepository) {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		c.Conn.Close()
	}()
	go repo.Subscribe(ctx, c.MsgChan, c.Chatbox)
	go func() {
		for {
			//receive message
			_, msg, err := c.Conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			mf := http.MessageFormat{
				Username: c.ClientName,
				Message:  string(msg),
			}
			if msg, err = json.Marshal(mf); err != nil {
				panic(err)
			}
			if err != nil {
				return
			}
			fmt.Printf("publish:%v:%v\n", c.ClientName, string(msg))
			repo.Publish(ctx, c.Chatbox, string(msg))
		}
	}()

	for {
		select {
		//send message
		case msg := <-c.MsgChan:
			c.Conn.WriteJSON(string(msg))
		}
	}

}
