package main

import (
	config "BlockApp/conf"
	_ "BlockApp/docs"
	"BlockApp/router"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// HelloWorld 示例接口
// @Summary 返回 Hello World
// @Description 这个接口返回 "Hello, World!"
// @Tags 示例
// @Produce json
// @Success 200 {string} string "Hello, World!"
// @Router /api/v1/hello [get]
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}

func main() {
	r := router.InitRouter()
	fmt.Println(config.Config.System)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//定义路由的GET方法及响应处理函数
	r.GET("/hello", HelloWorld)

	address := fmt.Sprintf("%s:%s", config.Config.System.AppHost, config.Config.System.AppPort)
	fmt.Println("✅ 服务器启动:", address)
	r.Run(address)
}
