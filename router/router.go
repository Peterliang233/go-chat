package router

import (
	"github.com/Peterliang233/go-chat/config"
	"github.com/Peterliang233/go-chat/middlerware"
	"github.com/Peterliang233/go-chat/router/v1/user"
	"github.com/Peterliang233/go-chat/service/socket"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化相关路由
func InitRouter() *gin.Engine {
	gin.SetMode(config.ServerSetting.AppMode)

	router := gin.New()

	router.Use(middlerware.Cors())
	router.Use(middlerware.Logger())
	chat := router.Group("/api/chat")

	hub := socket.NewHub()

	chat.Use(middlerware.JWTAuthMiddleware())
	{
		router.GET("/ws", func(c *gin.Context) {
			socket.ServerWs(hub, c)
		})
	}

	router.POST("/sign_up", user.Registry)
	router.POST("/sign_in", user.AuthHandler)

	go hub.Run()

	return router
}
