package handler

import (
	"expense-tracker/internal/response"
	"expense-tracker/internal/util/echocontext"
	requestmodel "expense-tracker/pkg/admin/model/request"
	"expense-tracker/pkg/admin/service"
	"github.com/labstack/echo/v4"
)

// StaffAuthInterface ...
type StaffAuthInterface interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

// StaffAuthImpl ...
type StaffAuthImpl struct{}

// StaffAuth ...
func StaffAuth() StaffAuthInterface {
	return StaffAuthImpl{}
}

// Register godoc
// @tags Auth
// @summary Register
// @id admin-staff-register
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.RegisterPayload true "Payload"
// @router /auth/register [POST]
func (StaffAuthImpl) Register(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.RegisterPayload)
		s       = service.StaffAuth()
	)

	token, err := s.Register(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, echo.Map{"token": token}, "")
}

// Login godoc
// @tags Auth
// @summary Login
// @id admin-staff-login
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.LoginPayload true "Payload"
// @router /auth/login [POST]
func (StaffAuthImpl) Login(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.LoginPayload)
		s       = service.StaffAuth()
	)

	token, err := s.Login(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, echo.Map{"token": token}, "")
}
