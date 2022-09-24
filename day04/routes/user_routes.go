package routes

import (
	"day02/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(router *echo.Echo, auth echo.MiddlewareFunc) {

	router.GET("/users", controllers.GetAllUsers(), auth)
	router.GET("/users/:id", controllers.GetUser(), auth)
	router.POST("/users", controllers.CreateUser())
	router.PUT("/users/:id", controllers.UpdateUser(), auth)
	router.DELETE("/users/:id", controllers.DeleteUser(), auth)
}
