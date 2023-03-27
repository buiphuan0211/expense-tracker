package routeauth

import (
	"expense-tracker/internal/auth"
	"expense-tracker/internal/response"
	"expense-tracker/internal/util/echocontext"
	"fmt"
	"github.com/labstack/echo/v4"
)

// RequiredLogin ...
func RequiredLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check invalid token
		user := auth.GetCurrentUserByToken(echocontext.GetToken(c))
		if user == nil || user.ID == "" {
			return response.R403(c, nil, response.CommonForbidden)
		}

		// Set Staff
		staff := &auth.User{
			ID:   user.ID,
			Name: user.Name,
		}

		fmt.Println("staff: ", staff) // DEL-LOG
		echocontext.SetStaff(c, staff)
		return next(c)
	}
}
