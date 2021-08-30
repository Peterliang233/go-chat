package router

import (
	"github.com/Peterliang233/go-chat/config"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化相关路由
func InitRouter() *gin.Engine {

	gin.SetMode(config.AppMode)

	router := gin.New()

	router.Use()
	{
		router.POST("/chat")
	}

	return router
}
