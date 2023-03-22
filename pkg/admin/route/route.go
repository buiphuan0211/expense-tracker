package route

import "github.com/labstack/echo/v4"

func Init(e *echo.Echo) {
	r := e.Group("/admin/expense")

	common(r)
}
