package handler

import (
	"CurlARC/internal/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, userHandler UserHandler) {

	e.POST("/signup", userHandler.SignUp())
	e.POST("/signin", userHandler.SignIn())
	e.GET("/users", userHandler.GetAllUser())

	// 認証が必要なルートにミドルウェアを適用
	authGroup := e.Group("/auth")
	authGroup.Use(middleware.JWTMiddleware)
	authGroup.GET("/me", userHandler.GetUser())
	authGroup.PATCH("/me", userHandler.UpdateUser())
	authGroup.DELETE("/me", userHandler.DeleteUser())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy!")
	})
}
