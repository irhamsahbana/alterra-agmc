package main

import (
	"day02/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	router := echo.New()

	routes.BookRoutes(router)
	routes.UserRoutes(router)

	router.Logger.Fatal(router.Start(":8080"))
}
