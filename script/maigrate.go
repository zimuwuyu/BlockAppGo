package main

import (
	config "BlockApp/conf"
	"BlockApp/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	pgsqlConfig := config.Config.Pgsql
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		pgsqlConfig.Host, pgsqlConfig.Username, pgsqlConfig.Password, pgsqlConfig.Database, pgsqlConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移：根据结构体生成数据表
	if err := db.AutoMigrate(
		&model.User{},
		model.BlockModel{},
		model.BlockModelType{},
		model.PictureStorage{},
		model.Task{},
		model.TaskLog{},
		model.UserFeedback{},
		model.CasbinRule{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 确认表已创建
	fmt.Println("Table created successfully")
}
