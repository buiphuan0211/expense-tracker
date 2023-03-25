package route

import (
	"expense-tracker/internal/util/routemiddleware"
	routeauth "expense-tracker/pkg/admin/route/auth"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.Use(routemiddleware.CORSConfig())

	e.Use(routeauth.Jwt())

	r := e.Group("/admin/expense")
	common(r)
	category(r)
	staff(r)
}
