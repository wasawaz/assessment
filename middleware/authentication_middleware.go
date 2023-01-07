// Package expense_middleware contains http handle on routes requests
package expense_middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// validAuth - is used for valid authentication value
const validAuth = "November 10, 2009"

// AuthMiddleware - is used for getting header Authorization and validate
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

// isValidAuthValue is used for checking value match specific string
func isValidAuthValue(value string) bool {
	return value == validAuth
}
