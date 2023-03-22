package echocontext

import (
	"github.com/labstack/echo/v4"
	"strings"
)

// GetToken ...
func GetToken(c echo.Context) string {
	token := GetHeaders(c).Get(echo.HeaderAuthorization)

	split := strings.Split(token, " ")
	if len(split) == 2 {
		return split[1]
	}

	return ""
}

// GetDeviceID ...
func GetDeviceID(c echo.Context) string {
	return GetHeaders(c).Get(HeaderDeviceID)
}
