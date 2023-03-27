package route

import (
	"expense-tracker/internal/auth/permission"
	"expense-tracker/pkg/admin/handler"
	routeauth "expense-tracker/pkg/admin/route/auth"
	routevalidation "expense-tracker/pkg/admin/route/validation"
	"github.com/labstack/echo/v4"
)

// category ...
func category(e *echo.Group) {
	var (
		g = e.Group("/categories", routeauth.RequiredLogin)
		h = handler.Category()
		v = routevalidation.Category()
	)

	view := routeauth.CheckPermission(permission.Scope.Category.View)

	// Create
	g.POST("", h.Create, v.Create)

	// All
	g.GET("", h.All, view, v.All)

	// Detail
	g.GET("/:id", h.Detail, v.ID)

	// Update
	g.PUT("/:id", h.Update, v.Update, v.ID)
}
