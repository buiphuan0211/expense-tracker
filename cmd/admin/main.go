package main

import (
	admin "expense-tracker/docs/admin"
	"expense-tracker/internal/config"
	"expense-tracker/internal/constant"
	"expense-tracker/pkg/admin/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
)

// @title manage-expenses - Admin API
// @version 1.0
// @description All APIs for expenses admin.
// @description
// @description ******************************
// @description - Add description
// @description ******************************
// @description
// @termsOfService
// @contact.name Dev An Bui
// @contact.url https://github.com/buiphuan0211
// @contact.email buiphuan0211@gmail.com
// @basePath /admin/expense

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()

	//e.Use(apmechov4.Middleware())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Gzip())
	if os.Getenv("ENV") == "release" {
		e.Use(middleware.Recover())
	}

	// Boostrap things
	server.Boostrap(e)

	if config.IsEnvDevelop() {
		domain := os.Getenv(constant.DomainAdmin)
		admin.SwaggerInfo.Host = domain
	}
	e.GET(admin.SwaggerInfo.BasePath+"/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(config.GetENV().Admin.Port))
}
