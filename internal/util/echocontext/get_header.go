package echocontext

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetHeaders ...
func GetHeaders(c echo.Context) http.Header {
	return c.Request().Header
}
