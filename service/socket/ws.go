package socket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is return msg
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager define a ws server manager
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[string]*Client),
}

// Start is  项目运行前, 协程开启start -> go Manager.Start()
func (manager *ClientManager) Start() {
	for {
		log.Println("<---管道通信--->")
		select {
		case conn := <-Manager.Register:
			log.Printf("新用户加入:%v", conn.ID)
			Manager.Clients[conn.ID] = conn
			jsonMessage, _ := json.Marshal(&Message{Content: "Successful connection to socket service"})
			conn.Send <- jsonMessage
		case conn := <-Manager.Unregister:
			log.Printf("用户离开:%v", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				jsonMessage, _ := json.Marshal(&Message{Content: "A socket has disconnected"})
				conn.Send <- jsonMessage
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case message := <-Manager.Broadcast:
			MessageStruct := Message{}
			json.Unmarshal(message, &MessageStruct)
			for id, conn := range Manager.Clients {
				if id != CreatId(MessageStruct.Recipient, MessageStruct.Sender) {
					continue
				}
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
		}
	}
}

// CreatId 创建一个用户id
func CreatId(uid, touid string) string {
	return uid + "_" + touid
}

// Read 读取消息
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		log.Printf("读取到客户端的信息:%s", string(message))
		Manager.Broadcast <- message
	}
}

// Write 写入消息
func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Printf("发送到到客户端的信息:%s", string(message))

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
