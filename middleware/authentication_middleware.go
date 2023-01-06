package expense_middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const validAuth = "November 10, 2009"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the "Authorization" header value
		authHeader := c.Request().Header.Get("Authorization")
		// Validate the value of the "Authorization" header
		if !isValidAuthValue(authHeader) {
			return c.NoContent(http.StatusUnauthorized)
		}

		// If the "Authorization" header is set and valid, call the next handler
		return next(c)
	}
}

func isValidAuthValue(value string) bool {
	return value == validAuth
}
