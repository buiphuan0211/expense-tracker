package route

import (
	routeauth "expense-tracker/pkg/admin/route/auth"
	"expense-tracker/pkg/admin/service"
	"github.com/labstack/echo/v4"
)

func migrate(e *echo.Group) {
	var g = e.Group("/migration", routeauth.RequiredLogin)
	g.GET("/insert-staff", func(c echo.Context) error {
		var s = service.Migrate()
		go s.MigrationStaffs()
		return c.JSON(200, "ok")
	})
}
