package handler

import (
	"CurlARC/internal/middleware"
	"net/http"

	_ "CurlARC/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRouting(
	e *echo.Echo,
	userHandler UserHandler,
	teamHandler TeamHandler,
	recordHandler RecordHandler,
) {
	// health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	// 認証が不要なエンドポイント
	// e.POST("/signup", userHandler.SignUp())
	e.POST("/authorize", userHandler.Authorize())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// 認証が必要なルートにミドルウェアを適用
	authGroup := e.Group("/auth")
	authGroup.Use(middleware.JWTMiddleware)

	// ユーザー関連のエンドポイント
	userGroup := authGroup.Group("/users")
	userGroup.GET("/me", userHandler.GetUser())
	userGroup.PATCH("/me", userHandler.UpdateUser())
	userGroup.DELETE("/me", userHandler.DeleteUser())
	userGroup.GET("/me/teams", teamHandler.GetTeamsByUserId())
	userGroup.GET("/me/teams/invited", teamHandler.GetInvitedTeams())

	// チーム関連のエンドポイント
	teamGroup := authGroup.Group("/teams")
	teamGroup.POST("/", teamHandler.CreateTeam())
	teamGroup.GET("/", teamHandler.GetAllTeams())
	teamGroup.GET("/:teamId/members", teamHandler.GetMembers())
	teamGroup.GET("/:teamId/details", teamHandler.GetTeamDetails())
	teamGroup.PATCH("/:teamId", teamHandler.UpdateTeam())
	teamGroup.DELETE("/:teamId", teamHandler.DeleteTeam())
	teamGroup.POST("/:teamId/invite", teamHandler.InviteUsers())
	teamGroup.POST("/:teamId/accept", teamHandler.AcceptInvitation())
	teamGroup.DELETE("/:teamId/:userId", teamHandler.RemoveMember())

	// レコード関連のエンドポイント
	recordGroup := authGroup.Group("/records")
	recordGroup.POST("/:teamId", recordHandler.CreateRecord())
	recordGroup.PATCH("/:recordId/append", recordHandler.AppendEndData())
	recordGroup.GET("/:recordId/details", recordHandler.GetRecordDetailsByRecordId())
	recordGroup.GET("/:teamId", recordHandler.GetRecordsByTeamId())
	recordGroup.PATCH("/:recordId", recordHandler.UpdateRecord())
	recordGroup.DELETE("/:recordId", recordHandler.DeleteRecord())
	recordGroup.PATCH("/:recordId/userId/visibility", recordHandler.SetVisibility())

	// デバッグ用
	debug := e.Group("/debug")
	debug.GET("/users", userHandler.GetAllUsers())
	debug.POST("/teams", teamHandler.CreateTeam())
	debug.GET("/teams", teamHandler.GetAllTeams())
	debug.GET("/teams/:teamId", teamHandler.GetMembers())
	debug.PATCH("/teams/:teamId/:userId", teamHandler.AcceptInvitation())
	debug.DELETE("/teams/:teamId/:userId", teamHandler.RemoveMember())
}
