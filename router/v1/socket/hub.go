package socket

import (
	"fmt"
	"strings"
)

// Hub 相当于一个事物管理中心
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// 房间号 key:client value:房间号
	roomID map[*Client]string
}

// NewHub .实例化一个hub
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		roomID:     make(map[*Client]string),
	}
}

// Run .监听客户端
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true                 // 注册client端
			h.roomID[client] = string(client.roomID) // 给client端添加房间号
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				delete(h.roomID, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				// 使用“&”对message进行message切割 获取房间号
				// 向信息所属的房间内的所有client 内添加send
				// msg[0]为房间号 msg[1]为打印内容
				msg := strings.Split(string(message), "&")
				fmt.Println("msg[0] :" + msg[0])
				fmt.Println("msg[1]:" + msg[1])
				if string(client.roomID) == msg[0] {
					select {
					case client.send <- []byte(msg[1]):
					default:
						close(client.send)
						delete(h.clients, client)
						delete(h.roomID, client)
					}
				}
			}
		}
	}
}
