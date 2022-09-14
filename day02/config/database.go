package config

import (
	"day02/models/entity"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := map[string]string{
		"host":     "localhost",
		"port":     "3306",
		"user":     "root",
		"password": "",
		"dbname":   "agmc-day02",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["user"], config["password"], config["host"], config["port"], config["dbname"])

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initMigrations(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}
