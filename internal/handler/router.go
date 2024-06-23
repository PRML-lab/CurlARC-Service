package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, userHandler UserHandler) {

	e.POST("/users/signup", userHandler.SignUp)
	e.POST("/users/signin", userHandler.SignIn)
	e.GET("/users/me", userHandler.GetUser)
	e.PUT("/users/me", userHandler.UpdateUser)
	e.DELETE("/users/me", userHandler.DeleteUser)
}
