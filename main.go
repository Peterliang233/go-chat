package main

import (
	"github.com/Peterliang233/go-chat/config"
	"github.com/Peterliang233/go-chat/database"
	"github.com/Peterliang233/go-chat/router"
	"github.com/Peterliang233/go-chat/service/socket"
)

func main() {
	database.InitDatabase()

	r := router.InitRouter()

	err := r.Run(config.ServerSetting.HttpPort)

	hub := socket.NewHub()

	go hub.Run()

	if err != nil {
		panic(err)
	}
}
