package routes

import (
	"day02/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(router *echo.Echo) {
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/users/:id", controllers.GetUser())
	router.POST("/users", controllers.CreateUser())
	router.PUT("/users/:id", controllers.UpdateUser())
	router.DELETE("/users/:id", controllers.DeleteUser())
}
