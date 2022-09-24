package config

import (
	"day02/models/entity"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["user"], config["password"], config["host"], config["port"], config["dbname"])

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", App.DBUser, App.DBPassword, App.DBHost, App.DBPort, App.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initMigrations(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}
