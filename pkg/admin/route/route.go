package route

import (
	"expense-tracker/internal/util/routemiddleware"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.Use(routemiddleware.CORSConfig())

	r := e.Group("/admin/expense")
	common(r)
	category(r)
}
