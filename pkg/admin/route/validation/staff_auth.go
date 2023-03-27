package routevalidation

import (
	"expense-tracker/internal/response"
	"expense-tracker/internal/util/echocontext"
	requestmodel "expense-tracker/pkg/admin/model/request"
	"github.com/labstack/echo/v4"
)

// StaffAuthInterface ...
type StaffAuthInterface interface {
	Register(next echo.HandlerFunc) echo.HandlerFunc
	Login(next echo.HandlerFunc) echo.HandlerFunc
}

// StaffAuthImpl ...
type StaffAuthImpl struct{}

// StaffAuth ...
func StaffAuth() StaffAuthInterface {
	return StaffAuthImpl{}
}

// Register ...
func (StaffAuthImpl) Register(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.RegisterPayload

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}

// Login ...
func (StaffAuthImpl) Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.LoginPayload
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}
