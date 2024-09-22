package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"CurlARC/internal/handler/response"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var jwtKey = os.Getenv("JWT_SECRET")

type Claims struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

// ParseJWTは、JWTを解析してその署名を検証します。
func ParseJWT(tokenStr string, secretKey string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// 署名アルゴリズムが一致しているか確認
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

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
		claims, err := ParseJWT(tokenStr, jwtKey) // jwtKeyを使って署名検証
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
