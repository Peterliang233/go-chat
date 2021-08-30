package main

import (
	"github.com/Peterliang233/go-chat/config"
	"github.com/Peterliang233/go-chat/database"
	"github.com/Peterliang233/go-chat/router"
)

func main() {
	database.InitDatabase()

	r := router.InitRouter()

	err := r.Run(config.HttpPort)

	if err != nil {
		panic(err)
	}
}
