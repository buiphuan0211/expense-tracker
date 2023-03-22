package server

import (
	"expense-tracker/internal/config/plogger"
	"expense-tracker/pkg/admin/route"
	"expense-tracker/pkg/admin/server/initialize"
	"github.com/labstack/echo/v4"
)

func Boostrap(e *echo.Echo) {
	// Init logger ...
	plogger.Init("expense-tracker", "expense-tracker-admin")

	// Init modules
	initialize.Init()

	// Routes
	route.Init(e)
}
