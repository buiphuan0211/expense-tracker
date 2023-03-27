package route

import (
	"expense-tracker/pkg/admin/service"
	"github.com/labstack/echo/v4"
)

func migrate(e *echo.Group) {
	var g = e.Group("/migration")
	g.GET("/insert-staff-admin", func(c echo.Context) error {
		var s = service.Migrate()
		go s.MigrationStaffAdmin()
		return c.JSON(200, "ok")
	})
}
