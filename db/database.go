package db

import (
	config "BlockApp/conf"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var PgsqlDB *gorm.DB

func InitPgSql() {
	pgsqlConfig := config.Config.Pgsql
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		pgsqlConfig.Host, pgsqlConfig.Username, pgsqlConfig.Password, pgsqlConfig.Database, pgsqlConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	PgsqlDB = db
	log.Println("PostgreSQL connection established successfully")

}
