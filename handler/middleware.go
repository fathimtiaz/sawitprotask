package handler

import "github.com/labstack/echo/v4"

// Authenticate middleware adds a `Server` header to the response.
func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}
