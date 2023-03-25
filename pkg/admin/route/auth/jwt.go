package routeauth

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/util/echocontext"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	envVars = config.GetENV()
)

func Jwt() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			token := echocontext.GetToken(c)
			return token == ""
		},
		SigningKey: []byte(envVars.Auth.SecretKey),
	})
}
