package router

import (
	"github.com/Peterliang233/go-chat/config"
	"github.com/Peterliang233/go-chat/middlerware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化相关路由
func InitRouter() *gin.Engine {
	gin.SetMode(config.AppMode)

	router := gin.New()

	router.Use(middlerware.Cors())
	router.Use(middlerware.Cors())

	router.POST("/chat")

	return router
}
