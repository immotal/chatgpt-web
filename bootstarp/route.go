package bootstarp

import (
	"github.com/gin-gonic/gin"
	"github.com/immotal/chatgpt-web/routes"
	"sync"
)

var router *gin.Engine
var once sync.Once

func SetUpRoute() {
	once.Do(func() {
		router = gin.Default()
		routes.RegisterWebRoutes(router)
	})
}
