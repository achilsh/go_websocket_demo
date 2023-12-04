package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
type WebSocketServer struct {
	Handlers map[string] gin.HandlerFunc
}

func (s *WebSocketServer)Start(g *gin.Engine) {
	g.SetTrustedProxies(nil)
	for k, v := range s.Handlers {
		g.GET(k, v)
	}
}

type WebSocketClient struct {
	Conn *websocket.Conn
	Data chan []byte
	Id string 
	MsgType int 
}
func (c *WebSocketClient)GetId() string {
	return c.Id
}
func (c *WebSocketClient) Read() {
	defer func() {

	}()

	for {
		_, msg, e := c.Conn.ReadMessage()
		if e != nil {
			fmt.Println("read msg from websocket connet fail, e: ", e)
			break
		}
		c.Data <- msg
	}
}
func (c *WebSocketClient) Write() {
	for v := range c.Data {
		e := c.Conn.WriteMessage(c.MsgType, v)
		if e != nil {
			fmt.Println("write websocket fail, e: ", e)
			break
		}
	}
}
func NewWsClient(c *websocket.Conn, id string, mtype int) *WebSocketClient{
	r := &WebSocketClient {
		Conn: c,
		Data: make(chan []byte),
		Id : id,
		MsgType: mtype,
	}
	return r
}