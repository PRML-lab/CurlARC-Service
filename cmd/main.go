package main

import (
	"CurlARC/internal/handler"
	"CurlARC/internal/injector"
	"CurlARC/internal/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: utils.GetAllowOrigins(), // 許可するオリジンのリスト
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization},
	}))

	// environment variables
	utils.LoadEnv()

	// Handler
	userHandler := injector.InjectUserHandler()

	// Routing
	handler.InitRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
