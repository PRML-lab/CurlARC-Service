package middleware

import (
	"bytes"
	"io"

	"github.com/labstack/echo/v4"
)

func LogBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		buf, _ := io.ReadAll(c.Request().Body)
		c.Logger().Info(string(buf))
		c.Request().Body = io.NopCloser(bytes.NewBuffer(buf))
		return next(c)
	}
}