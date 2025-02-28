package main

import (
	"BlockApp/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=BlockApp port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移：根据结构体生成数据表
	if err := db.AutoMigrate(&model.User{}, model.BlockModel{}, model.BlockModelType{}, model.PictureStorage{}, model.Task{}, model.TaskLog{}, model.UserFeedback{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 确认表已创建
	fmt.Println("Table created successfully")
}
