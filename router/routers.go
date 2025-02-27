package router

import (
	config "BlockApp/conf"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Config.System.Env)
	router := gin.New()
	router.Use(gin.Recovery())
	register(router)
	return router
}

func register(router *gin.Engine) {
	InitBlockRoutes(router)
}
