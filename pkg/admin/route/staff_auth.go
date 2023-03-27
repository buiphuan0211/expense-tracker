package route

import (
	"expense-tracker/pkg/admin/handler"
	routevalidation "expense-tracker/pkg/admin/route/validation"
	"github.com/labstack/echo/v4"
)

func staffAuth(e *echo.Group) {
	var (
		g = e.Group("/auth")
		h = handler.StaffAuth()
		v = routevalidation.StaffAuth()
	)

	g.POST("/register", h.Register, v.Register)

	g.POST("/login", h.Login, v.Login)
}
