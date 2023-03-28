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
	UpdatePermissions(c echo.Context) error
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

// UpdatePermissions godoc
// @tags Staff
// @summary UpdatePermissions
// @id admin-staff-update-permissions
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Staff id"
// @param payload body requestmodel.StaffPermissionUpdatePayload true "Payload"
// @success 200 {object} responsemodel.Upsert
// @router /staffs/{id}/update-permissions [PATCH]
func (staffImpl) UpdatePermissions(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		staff   = echocontext.GetStaff(c).(*auth.User)
		payload = echocontext.GetPayload(c).(requestmodel.StaffPermissionUpdatePayload)
		id      = echocontext.GetParam(c, "id").(string)
		s       = service.Staff(staff)
	)

	res, err := s.UpdatePermissions(ctx, id, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}
