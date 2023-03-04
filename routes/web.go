package routes

import (
	"github.com/gin-gonic/gin"
	. "github.com/immotal/chatgpt-web/app/http/controllers"
	"github.com/immotal/chatgpt-web/app/middlewares"
)

var chatController = NewChatController()

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {
	router.Use(middlewares.Cors())
	router.GET("/", chatController.Index)
	router.POST("/completion", chatController.Completion)
}
