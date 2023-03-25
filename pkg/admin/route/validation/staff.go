package routevalidation

import (
	"github.com/labstack/echo/v4"
)

type StaffInterface interface {
	All(next echo.HandlerFunc) echo.HandlerFunc
}

type staffImpl struct{}

func Staff() StaffInterface {
	return staffImpl{}
}

func (staffImpl) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		return next(c)
	}
}
