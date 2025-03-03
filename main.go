package main

import (
	config "BlockApp/conf"
	"BlockApp/db"
	_ "BlockApp/docs"
	"BlockApp/router"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := router.InitRouter()
	fmt.Println(config.Config.System)
	db.InitPgSql()
	// 加载swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 拼接服务地址
	address := fmt.Sprintf("%s:%s", config.Config.System.AppHost, config.Config.System.AppPort)
	fmt.Println("✅ 服务器启动:", address)
	err := r.Run(address)
	if err != nil {
		return
	}
}
