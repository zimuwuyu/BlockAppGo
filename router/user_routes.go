package router

import (
	"BlockApp/controller"
	"github.com/gin-gonic/gin"
)

var (
	User = &controller.UserController{}
)

func InitUserRoutes(r *gin.Engine) gin.IRoutes {
	blockModel := r.Group("/v1")
	blockModel.POST("/login", User.UserLogin)
	return r
}
