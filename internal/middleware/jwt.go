package middleware

import (
	"net/http"

	"CurlARC/internal/utils"

	"github.com/labstack/echo"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
		}

		tokenStr := authHeader[len("Bearer "):]
		claims, err := utils.ParseJWT(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		c.Set("uid", claims.UID)
		return next(c)
	}
}
