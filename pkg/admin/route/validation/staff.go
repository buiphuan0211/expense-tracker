package routevalidation

import (
	"expense-tracker/internal/response"
	"expense-tracker/internal/util/echocontext"
	"expense-tracker/internal/util/pmongo"
	"expense-tracker/pkg/admin/errorcode"
	requestmodel "expense-tracker/pkg/admin/model/request"
	"github.com/labstack/echo/v4"
)

// StaffInterface ...
type StaffInterface interface {
	All(next echo.HandlerFunc) echo.HandlerFunc
	UpdatePermissions(next echo.HandlerFunc) echo.HandlerFunc
	ID(next echo.HandlerFunc) echo.HandlerFunc
}

// staffImpl ...
type staffImpl struct{}

// Staff ...
func Staff() StaffInterface {
	return staffImpl{}
}

// All ...
func (staffImpl) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query requestmodel.StaffAll

		if err := c.Bind(&query); err != nil {
			return response.R400(c, nil, "")
		}

		if err := query.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetQuery(c, query)
		return next(c)
	}
}

// UpdatePermissions ...
func (staffImpl) UpdatePermissions(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.StaffPermissionUpdatePayload
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

// ID ...
func (staffImpl) ID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = c.Param("id")

		if valid := pmongo.IsValidID(id); !valid {
			return response.R400(c, nil, errorcode.CategoryExistedName)
		}

		echocontext.SetParam(c, "id", id)
		return next(c)
	}
}
