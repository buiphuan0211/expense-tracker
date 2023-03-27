package routeauth

import (
	"expense-tracker/internal/auth"
	"expense-tracker/internal/response"
	"expense-tracker/internal/util/echocontext"
	"github.com/labstack/echo/v4"
)

// CheckPermission ...
func CheckPermission(scopes []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Check invalid token
			staff := auth.GetCurrentUserByToken(echocontext.GetToken(c))
			if staff == nil || staff.ID == "" {
				return response.R403(c, nil, response.CommonForbidden)
			}

			if isValid := auth.CheckPermission(scopes, echocontext.GetToken(c)); !isValid {
				return response.R401(c, nil, response.CommonUnauthorized)
			}

			return next(c)
		}
	}
}
