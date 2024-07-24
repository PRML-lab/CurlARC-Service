package middleware

import (
	"net/http"

	"CurlARC/internal/handler/response"
	"CurlARC/internal/utils"

	"github.com/labstack/echo"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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

		tokenStr := authHeader[len("Bearer "):]
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

		c.Set("uid", claims.UID)
		return next(c)
	}
}
