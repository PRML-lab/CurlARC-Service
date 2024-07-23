package handler

import (
	"CurlARC/internal/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, userHandler UserHandler, teamHandler TeamHandler) {

	e.POST("/signup", userHandler.SignUp())
	e.POST("/signin", userHandler.SignIn())

	// デバッグ用
	debug := e.Group("/debug")
	debug.GET("/users", userHandler.GetAllUsers())
	debug.POST("/teams", teamHandler.CreateTeam())
	debug.GET("/teams", teamHandler.GetAllTeams())
	debug.GET("/teams/:teamId", teamHandler.GetMembers())
	debug.POST("/teams/:teamId/:userId", teamHandler.InviteUser())
	debug.PATCH("/teams/:teamId/:userId", teamHandler.AcceptInvitation())
	debug.DELETE("/teams/:teamId/:userId", teamHandler.RemoveMember())

	// 認証が必要なルートにミドルウェアを適用
	authGroup := e.Group("/auth")
	authGroup.Use(middleware.JWTMiddleware)
	authGroup.GET("/me", userHandler.GetUser())
	authGroup.PATCH("/me", userHandler.UpdateUser())
	authGroup.DELETE("/me", userHandler.DeleteUser())

	//team
	authGroup.POST("/teams", teamHandler.CreateTeam())
	authGroup.GET("/teams", teamHandler.GetAllTeams())

	authGroup.GET("/teams/:teamId", teamHandler.GetMembers())
	authGroup.PATCH("/teams/:teamId", teamHandler.UpdateTeam())
	authGroup.DELETE("/teams/:teamId", teamHandler.DeleteTeam())

	authGroup.POST("/teams/:teamId/:userId", teamHandler.InviteUser())
	authGroup.PATCH("/teams/:teamId/:userId", teamHandler.AcceptInvitation())
	authGroup.DELETE("/teams/:teamId/:userId", teamHandler.RemoveMember())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy!")
	})
}
