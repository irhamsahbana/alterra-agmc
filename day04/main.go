package main

import (
	"day02/config"
	"day02/lib/middleware"
	"day02/routes"
	"errors"

	"github.com/labstack/echo/v4"
	middlewareEcho "github.com/labstack/echo/v4/middleware"
)

func main() {

	router := echo.New()
	router.Debug = true

	middleware.LogMiddleware(router)

	jwtSecret := config.App.JwtSecret
	auth := middlewareEcho.JWTWithConfig(middlewareEcho.JWTConfig{
		SigningKey: []byte(jwtSecret),
		ErrorHandler: func(err error) error {
			return errors.New("Unauthorized")
		},
	})

	routes.AuthRoutes(router)
	routes.BookRoutes(router, auth)
	routes.UserRoutes(router, auth)

	router.Logger.Fatal(router.Start(":8080"))
}
