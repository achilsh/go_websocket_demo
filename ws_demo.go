package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var RouteMap map[string] gin.HandlerFunc
var wsClients map[string] *WebSocketClient
func init () {
	RouteMap = make(map[string] gin.HandlerFunc)
	wsClients = make(map[string] *WebSocketClient)

	RouteMap["/test/v1"] = func(g *gin.Context) {
		fmt.Println("is connect....")
		upgrade := websocket.Upgrader {
			CheckOrigin: func(r *http.Request) bool { return true },
		 }

		 c, e := upgrade.Upgrade(g.Writer, g.Request, nil )
		 if e != nil {
			fmt.Println("update fail, e: ", e)
			return 
		 }
		 //new conn now. 
		 
		 cli := NewWsClient(c, uuid.New().String(), websocket.BinaryMessage)
		 wsClients[cli.GetId()] = cli

		 //
		 go cli.Read()
		 go cli.Write()
	}
}

func TimerSendData() {
	i := 0
	tmr := time.NewTicker(2*time.Second)
	for {
		select {
		case <-tmr.C:
			for _, v :=range wsClients {
				v.Data <- []byte("this is demo")
			}
			i++
			if i > 5 {
				break
			}
		}
	}
}
func main() {
	g := gin.Default()
	gin.SetMode(gin.DebugMode)

	wsSrv := &WebSocketServer{
		Handlers: RouteMap,
	}
	wsSrv.Start(g)

	go TimerSendData()
	g.Run("127.0.0.1:8090")
}