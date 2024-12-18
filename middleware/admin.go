package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var AdminCheck = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Request().Context().Value("role").(string)
		if role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "You are not allowed to access this resource")
		}
		return next(c)
	}
}
