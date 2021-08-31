package model

import "time"

// User 用户表
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Room 房间表
type Room struct {
	ID       int `json:"id"`
	OwnerID  int `json:"owner_id"`
	EnterKey int `json:"enter_key"`
}

// Message 消息表
type Message struct {
	ID       int        `json:"id"`
	OwnerID  int        `json:"owner_id"`
	RoomID   int        `json:"room_id"`
	SendTime *time.Time `json:"send_time"`
	Content  string     `json:"content"`
}
