package echocontext

import "github.com/labstack/echo/v4"

// GetStaffID ...
func GetStaffID(c echo.Context) string {
	staffIDInterface := c.Get(KeyStaffID)
	if staffIDInterface == nil {
		return ""
	}
	return staffIDInterface.(string)
}
