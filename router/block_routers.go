package router

import (
	"BlockApp/controller"
	"github.com/gin-gonic/gin"
)

var (
	Block = &controller.BlockController{}
)

func InitBlockRoutes(r *gin.Engine) gin.IRoutes {
	blockModel := r.Group("/blockModel")
	blockModel.GET("", Block.Get)
	return r
}
