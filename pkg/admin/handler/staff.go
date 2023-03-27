package handler

import (
	"expense-tracker/internal/auth"
	"expense-tracker/internal/response"
	"expense-tracker/internal/util/echocontext"
	"expense-tracker/internal/util/mgquery"
	requestmodel "expense-tracker/pkg/admin/model/request"
	"expense-tracker/pkg/admin/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type StaffInterface interface {
	All(c echo.Context) error
}

type staffImpl struct{}

func Staff() StaffInterface {
	return staffImpl{}
}

// All godoc
// @tags Staff
// @summary All
// @id admin-staff-all
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query requestmodel.StaffAll true "Query"
// @success 200 {object} responsemodel.StaffAll
// @router /staffs [GET]
func (staffImpl) All(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		qParams = echocontext.GetQuery(c).(requestmodel.StaffAll)
		staff   = echocontext.GetStaff(c).(*auth.User)
		s       = service.Staff(staff)
	)
	q := mgquery.AppQuery{
		Page:          qParams.Page,
		Limit:         qParams.Limit,
		Keyword:       qParams.Keyword,
		Status:        qParams.Status,
		SortInterface: bson.M{"createdAt": -1},
	}

	res := s.All(ctx, q)
	return response.R200(c, res, "")
}
