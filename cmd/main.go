package main

import (
	"CurlARC/internal/handler"
	"CurlARC/internal/injector"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // 許可するオリジンのリスト
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization},
	}))

	// environment variables
	loadEnv()

	// Handler
	userHandler := injector.InjectUserHandler()

	// Routing
	handler.InitRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
