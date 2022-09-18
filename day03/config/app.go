package config

import "gorm.io/gorm"

var App *Application

type Application struct {
	Sql        *gorm.DB
	JwtSecret  string
	DBPort     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

func init() {
	InitApp()
}

func InitApp() {
	App = &Application{}
	App.DBPort = "3306"
	App.DBHost = "localhost"
	App.DBUser = "root"
	App.DBPassword = ""
	App.DBName = "agmc-day02"

	App.Sql = InitDB()
	initMigrations(App.Sql)

	App.JwtSecret = "secret"
}
