package router

import (
	"github.com/Peterliang233/go-chat/config"
	"github.com/Peterliang233/go-chat/middlerware"
	"github.com/Peterliang233/go-chat/router/v1/socket"
	"github.com/Peterliang233/go-chat/router/v1/user"
	Service "github.com/Peterliang233/go-chat/service/socket"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化相关路由
func InitRouter() *gin.Engine {
	gin.SetMode(config.ServerSetting.AppMode)

	router := gin.New()

	router.Use(middlerware.Cors())

	router.Use(middlerware.Logger())

	go Service.Manager.Start()

	router.GET("/ws", socket.WsHandler)

	router.POST("/sign_up", user.Registry)

	router.POST("/sign_in", user.AuthHandler)

	return router
}
