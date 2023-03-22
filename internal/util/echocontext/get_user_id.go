package echocontext

import "github.com/labstack/echo/v4"

// GetUserID ...
func GetUserID(c echo.Context) string {
	userIDInterface := c.Get(KeyUserID)
	if userIDInterface == nil {
		return ""
	}
	return userIDInterface.(string)
}
