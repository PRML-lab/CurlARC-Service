package middleware

import (
	"net/http"
	"strings"

	"CurlARC/internal/handler/response"
	"CurlARC/internal/utils"

	"github.com/labstack/echo/v4"
)

// JWTMiddlewareはリクエストのAuthorizationヘッダーからJWTを取得し、検証します。
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Authorizationヘッダーからトークンを取得
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusUnauthorized,
					Message: "missing token",
				},
			})
		}

		// Bearerスキームを確認し、トークン部分を抽出
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusUnauthorized,
					Message: "invalid token format",
				},
			})
		}

		tokenStr := bearerToken[1]

		// JWTトークンの解析と検証
		claims, err := utils.ParseJWT(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusUnauthorized,
					Message: "invalid token",
				},
			})
		}

		// ユーザーIDをコンテキストに設定
		c.Set("uid", claims.UID)
		return next(c)
	}
}
