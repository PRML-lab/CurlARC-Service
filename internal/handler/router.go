package handler

import (
	"CurlARC/internal/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, userHandler UserHandler, teamHandler TeamHandler) {

	e.POST("/signup", userHandler.SignUp())
	e.POST("/signin", userHandler.SignIn())
	e.GET("/users", userHandler.GetAllUser())

	// 認証が必要なルートにミドルウェアを適用
	authGroup := e.Group("/auth")
	authGroup.Use(middleware.JWTMiddleware)
	authGroup.GET("/me", userHandler.GetUser())
	authGroup.PATCH("/me", userHandler.UpdateUser())
	authGroup.DELETE("/me", userHandler.DeleteUser())

	//team
	authGroup.POST("/teams", teamHandler.CreateTeam())
	authGroup.GET("/teams", teamHandler.GetAllTeams())
	authGroup.GET("/teams/:id", teamHandler.GetTeam())
	authGroup.PATCH("/teams/:id", teamHandler.UpdateTeam())
	authGroup.DELETE("/teams/:id", teamHandler.DeleteTeam())
	authGroup.POST("/teams/:teamId/:userId", teamHandler.AddMember())
	authGroup.DELETE("/teams/:teamId/:userId", teamHandler.RemoveMember())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy!")
	})
}
