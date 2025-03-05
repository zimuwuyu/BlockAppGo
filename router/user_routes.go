package router

import (
	"BlockApp/controller"
	"github.com/gin-gonic/gin"
)

var (
	User = &controller.UserController{}
)

func InitUserRoutes(r *gin.Engine) gin.IRoutes {
	UserModel := r.Group("/v1")
	UserModel.POST("/", User.Create)
	UserModel.POST("/login", User.UserLogin)
	UserModel.POST("/refreshToken", User.RefreshToken)
	//UserModel.POST("/register", User.UserRegister)
	return r
}
