package routevalidation

import (
	"expense-tracker/internal/response"
	"expense-tracker/internal/util/echocontext"
	requestmodel "expense-tracker/pkg/admin/model/request"
	"github.com/labstack/echo/v4"
)

// StaffInterface ...
type StaffInterface interface {
	All(next echo.HandlerFunc) echo.HandlerFunc
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
