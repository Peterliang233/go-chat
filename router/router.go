package router

import (
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化相关路由
func InitRouter() *gin.Engine {

	router := gin.New()

	return router
}
