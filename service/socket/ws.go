package socket

import (
	"bytes"
	"github.com/Peterliang233/go-chat/config"
	"github.com/Peterliang233/go-chat/database"
	"github.com/Peterliang233/go-chat/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"time"
)

var (
	writeWait      = config.SocketSetting.WriteWait
	pongWait       = config.SocketSetting.PongWait
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = config.SocketSetting.MaxMessageSize
)

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  config.SocketSetting.ReadBufferSize,
	WriteBufferSize: config.SocketSetting.WriteBufferSize,
}

type Client struct {
	hub      *Hub
	coon     *websocket.Conn
	send     chan []byte
	username []byte
	roomID   []byte
}

// ReadPump 从消息中心读取信息
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		_ = c.coon.Close()
	}()

	c.coon.SetReadLimit(maxMessageSize)
	_ = c.coon.SetReadDeadline(time.Now().Add(pongWait))
	c.coon.SetPongHandler(func(string) error {
		_ = c.coon.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.coon.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Fatalf("Read Message error,%v", err)
			}

			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))

		c.hub.broadcast <- message
	}
}

// WritePump 向消息中心写入信息
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()

		_ = c.coon.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			_ = c.coon.SetWriteDeadline(time.Now().Add(writeWait))
			if ok {
				_ = c.coon.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.coon.NextWriter(websocket.TextMessage)

			if err != nil {
				return
			}

			_, _ = w.Write(message)

			n := len(message)

			for i := 0; i < n; i++ {
				_, _ = w.Write(newLine)
				_, _ = w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				log.Printf("error,%v", err)
				return
			}
		case <-ticker.C:
			_ = c.coon.SetWriteDeadline(time.Now().Add(writeWait))

			if err := c.coon.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ServerWs 开启socket通信
func ServerWs(hub *Hub, c *gin.Context) {
	var chat model.Room

	_ = c.ShouldBind(&chat)

	roomID := strconv.Itoa(chat.ID)
	username, err := GetUsernameByID(chat.OwnerID)

	if err != nil {
		log.Fatalf("err %v", err)
	}

	var upgrade websocket.Upgrader

	coon, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("err %v", err)
		return
	}

	client := &Client{
		hub:      hub,
		coon:     coon,
		send:     make(chan []byte, 256),
		username: []byte(username),
		roomID:   []byte(roomID),
	}

	client.hub.register <- client

	go client.ReadPump()
	go client.WritePump()
}

// GetUsernameByID 通过用户id获取用户名
func GetUsernameByID(id int) (username string, err error) {
	var u model.User

	if err := database.Db.
		Where("id = ?", id).
		First(&u).
		Error; err != nil {
		return "", err
	}

	return u.Username, nil
}
