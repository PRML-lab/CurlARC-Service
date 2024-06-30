package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, userHandler UserHandler) {

	e.POST("/signup", userHandler.SignUp())
	e.POST("/signin", userHandler.SignIn)

	// 認証が必要なルートにミドルウェアを適用
	authGroup := e.Group("/auth")
	authGroup.Use(userHandler.Authenticate)
	authGroup.GET("/user", userHandler.GetUser)
	authGroup.PATCH("/user", userHandler.UpdateUser)
	authGroup.DELETE("/user", userHandler.DeleteUser)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy!")
	})
}
