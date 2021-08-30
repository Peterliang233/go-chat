package main

import (
	"github.com/Peterliang233/go-chat/database"
	router2 "github.com/Peterliang233/go-chat/router"
)

func main() {
	database.InitDatabase()

	router := router2.InitRouter()

}
