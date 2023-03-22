package route

import (
	"expense-tracker/pkg/admin/handler"
	routevalidation "expense-tracker/pkg/admin/route/validation"
	"github.com/labstack/echo/v4"
)

// category ...
func category(e *echo.Group) {
	var (
		g = e.Group("/categories")
		h = handler.Category()
		v = routevalidation.Category()
	)

	g.POST("", h.Create, v.Create)
	g.GET("", h.All, v.All)
}
