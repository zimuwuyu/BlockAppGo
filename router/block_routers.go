package router

import (
	"BlockApp/controller"
	"BlockApp/middleware"
	"github.com/gin-gonic/gin"
)

var (
	Block = &controller.BlockController{}
)

func InitBlockRoutes(r *gin.Engine) gin.IRoutes {
	blockModel := r.Group("/v1")
	blockModel.Use(middleware.JWTAuthMiddleware())
	blockModel.GET("/blockModel", Block.GetBlocBModel)
	return r
}
