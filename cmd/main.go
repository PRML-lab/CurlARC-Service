package main

import (
	"CurlARC/internal/handler"
	"CurlARC/internal/injector"

	"github.com/labstack/echo"
)

func main() {

	userHandler := injector.InjectUserHandler()
	e := echo.New()
	handler.InitRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
