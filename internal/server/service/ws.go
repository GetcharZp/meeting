package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

// roomIdentity => { userIdentity => wsConn, ... }
var wsP2pConnMap = sync.Map{}

func WsP2PConnection(c *gin.Context) {
	// 0. 获取房间和用户的信息
	// 1. 升级协议
	// 2. 存储当前的连接信息
	// 3. 监听发过来的消息

	// 获取房间和用户的信息
	in := new(WsP2PConnectionRequest)
	err := c.ShouldBindUri(in)
	if err != nil {
		log.Println("ShouldBindUri err.", err)
		return
	}

	// 升级协议
	var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	// 存储当前的连接信息
	var userConnMap = new(sync.Map)
	value, ok := wsP2pConnMap.Load(in.RoomIdentity)
	if ok {
		userConnMap = value.(*sync.Map)
	}
	userConnMap.Store(in.UserIdentity, conn)
	wsP2pConnMap.Store(in.RoomIdentity, userConnMap)

	// 监听发过来的消息
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			return
		}
		v, okk := wsP2pConnMap.Load(in.RoomIdentity)
		if okk {
			v.(*sync.Map).Range(func(key, value interface{}) bool {
				err := value.(*websocket.Conn).WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Println("WriteMessage err.", err)
				}
				return true
			})
		}
	}
}
