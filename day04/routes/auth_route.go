package routes

import (
	"day02/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(router *echo.Echo) {
	router.POST("/login", controllers.Login())
	router.POST("/register", controllers.Register())
}
