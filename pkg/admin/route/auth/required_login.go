package routeauth

import (
	"expense-tracker/internal/auth"
	"expense-tracker/internal/util/echocontext"
	"github.com/labstack/echo/v4"
)

// RequiredLogin ...
func RequiredLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check invalid token
		//user := auth.GetCurrentUserByToken(c.Get("user"))
		//if user == nil || user.ID == "" {
		//	return response.R403(c, nil, response.CommonForbidden)
		//}

		//// Set Staff
		//staff := &auth.User{
		//	ID:   user.ID,
		//	Name: user.Name,
		//}

		staff := &auth.User{
			ID:   "641c857fdf327cd4ebd56de0",
			Name: "Bui Phu An",
		}

		echocontext.SetStaff(c, staff)
		return next(c)
	}
}
