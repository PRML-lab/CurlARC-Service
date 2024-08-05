package middleware

import (
	"net/http"

	"CurlARC/internal/handler/response"
	"CurlARC/internal/utils"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusUnauthorized,
					Message: "missing token",
				},
			})
		}

		tokenStr := cookie.Value
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
