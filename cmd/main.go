package main

import (
	"CurlARC/internal/handler"
	"CurlARC/internal/injector"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Handler
	userHandler := injector.InjectUserHandler()

	// Routing
	handler.InitRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
