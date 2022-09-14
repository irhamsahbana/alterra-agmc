package config

import "gorm.io/gorm"

var App *Application

type Application struct {
	Sql *gorm.DB
}

func init() {
	InitApp()
}

func InitApp() {
	App = &Application{}
	App.Sql = InitDB()

	initMigrations(App.Sql)
}
