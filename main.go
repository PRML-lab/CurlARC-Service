// @title CurlARC API
// @version 1.0
package main

import (
	"CurlARC/internal/handler"
	"CurlARC/internal/injector"
	"CurlARC/internal/utils"

	myMiddleware "CurlARC/internal/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// environment variables
	utils.LoadEnv()

	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/health"
		},
	}))
	e.Use(myMiddleware.LogBody)

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     utils.GetAllowOrigins(), // 許可するオリジンのリスト
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	// Handler
	userHandler := injector.InjectUserHandler()
	recordHandler := injector.InjectRecordHandler()
	teamHandler := injector.InjectTeamHandler()

	// Routing
	handler.InitRouting(e, userHandler, teamHandler, recordHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
