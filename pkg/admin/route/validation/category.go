package routevalidation

import (
	"expense-tracker/internal/response"
	"expense-tracker/internal/util/echocontext"
	requestmodel "expense-tracker/pkg/admin/model/request"
	"github.com/labstack/echo/v4"
)

// CategoryInterface ...
type CategoryInterface interface {
	Create(next echo.HandlerFunc) echo.HandlerFunc
	All(next echo.HandlerFunc) echo.HandlerFunc
}

// categoryImpl ...
type categoryImpl struct{}

// Category ...
func Category() CategoryInterface {
	return categoryImpl{}
}

// Create ...
func (categoryImpl) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.CategoryCreatePayload

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

// All ...
func (categoryImpl) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query requestmodel.CategoryAll
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
