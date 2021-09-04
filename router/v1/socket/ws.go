package socket

import (
	Service "github.com/Peterliang233/go-chat/service/socket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// WsHandler socket通信相关接口，对http协议升级为socket相关协议
func WsHandler(c *gin.Context) {
	uid := c.Query("uid")
	touid := c.Query("to_uid")
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// 可以添加用户信息验证
	client := &Service.Client{
		ID:     Service.CreatId(uid, touid),
		Socket: conn,
		Send:   make(chan []byte),
	}
	Service.Manager.Register <- client

	go client.Read()

	go client.Write()
}
