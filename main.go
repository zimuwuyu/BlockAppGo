// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @tokenUrl /v1/login
// @in header
// @name Authorization
package main

import (
	config "BlockApp/conf"
	"BlockApp/db"
	_ "BlockApp/docs"
	"BlockApp/middleware"
	"BlockApp/router"
	"context"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	r := router.InitRouter()
	fmt.Println(config.Config.System)
	db.InitPgSql()
	pgsqlConfig := config.Config.Pgsql
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		pgsqlConfig.Host, pgsqlConfig.Username, pgsqlConfig.Password, pgsqlConfig.Database, pgsqlConfig.Port)
	casbinSrv := middleware.NewCasbinService(dsn)
	r.Use(middleware.CasbinMiddleware(casbinSrv))
	// 加载swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 拼接服务地址
	address := fmt.Sprintf("%s:%s", config.Config.System.AppHost, config.Config.System.AppPort)
	srv := &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  config.Config.System.ReadTimeOut,
		WriteTimeout: config.Config.System.WriteTimeout,
	}
	go func() {
		fmt.Printf("✅ swagger address http://%s/swagger/index.html", address)
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
