package routes

import (
	. "giithub.com/immotal/chatgpt-web/app/http/controllers"
	"giithub.com/immotal/chatgpt-web/app/middlewares"
	"github.com/gin-gonic/gin"
)

var chatController = NewChatController()

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {
	router.Use(middlewares.Cors())
	router.GET("/", chatController.Index)
	router.POST("/completion", chatController.Completion)
}
