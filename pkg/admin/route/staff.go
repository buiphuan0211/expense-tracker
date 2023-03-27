package route

import (
	"expense-tracker/pkg/admin/handler"
	routeauth "expense-tracker/pkg/admin/route/auth"
	routevalidation "expense-tracker/pkg/admin/route/validation"
	"github.com/labstack/echo/v4"
)

func staff(e *echo.Group) {
	var (
		g = e.Group("/staffs", routeauth.RequiredLogin)
		h = handler.Staff()
		v = routevalidation.Staff()
	)

	// All
	g.GET("", h.All, v.All)
}
