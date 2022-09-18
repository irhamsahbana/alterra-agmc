package routes

import (
	"day02/controllers"

	"github.com/labstack/echo/v4"
)

func BookRoutes(router *echo.Echo, auth echo.MiddlewareFunc) {

	router.GET("/books", controllers.GetAllBooks())
	router.GET("/books/:id", controllers.GetBook())

	router.POST("/books", controllers.CreateBook(), auth)
	router.PUT("/books/:id", controllers.UpdateBook(), auth)
	router.DELETE("/books/:id", controllers.DeleteBook(), auth)
}
