package router

import (
	"github.com/Peterliang233/go-chat/config"
	"github.com/Peterliang233/go-chat/middlerware"
	"github.com/Peterliang233/go-chat/router/v1/user"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化相关路由
func InitRouter() *gin.Engine {
	gin.SetMode(config.ServerSetting.AppMode)

	router := gin.New()

	router.Use(middlerware.Cors())
	router.Use(middlerware.Logger())

	router.POST("/sign_up", user.Registry)
	router.POST("/sign_in", user.Login)

	return router
}
