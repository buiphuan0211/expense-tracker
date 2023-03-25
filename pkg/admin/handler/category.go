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

// CategoryInterface ...
type CategoryInterface interface {
	Create(c echo.Context) error
	All(c echo.Context) error
	Detail(c echo.Context) error
	Update(c echo.Context) error
}

// categoryImpl ...
type categoryImpl struct{}

// Category ...
func Category() CategoryInterface {
	return categoryImpl{}
}

// Create doc
// @tags Category
// @summary Create
// @id admin-category-create
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.CategoryCreatePayload true "Payload"
// @success 200 {object} responsemodel.Upsert
// @router /categories [post]
func (categoryImpl) Create(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.CategoryCreatePayload)
		staff   = echocontext.GetStaff(c).(*auth.User)
		s       = service.Category(staff)
	)

	res, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

// All godoc
// @tags Category
// @summary All
// @id admin-category-all
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query requestmodel.CategoryAll true "Query"
// @success 200 {object} responsemodel.CategoryAll
// @router /categories [GET]
func (categoryImpl) All(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		qParams = echocontext.GetQuery(c).(requestmodel.CategoryAll)
		staff   = echocontext.GetStaff(c).(*auth.User)
		s       = service.Category(staff)
	)

	q := mgquery.AppQuery{
		Page:          qParams.Page,
		Limit:         qParams.Limit,
		Keyword:       qParams.Keyword,
		SortInterface: bson.M{"createdAt": -1},
	}

	res := s.All(ctx, q)
	return response.R200(c, res, "")
}

// Detail godoc
// @tags Category
// @summary Detail
// @id admin-category-detail
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Category id"
// @success 200 {object} responsemodel.CategoryDetail
// @router /categories/{id} [GET]
func (categoryImpl) Detail(c echo.Context) error {
	var (
		ctx   = echocontext.GetContext(c)
		id    = echocontext.GetParam(c, "id").(string)
		staff = echocontext.GetStaff(c).(*auth.User)
		s     = service.Category(staff)
	)

	res, err := s.Detail(ctx, id)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

// Update doc
// @tags Category
// @summary Update
// @id admin-category-update
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Category id"
// @param payload body requestmodel.CategoryUpdatePayload true "Payload"
// @success 200 {object} responsemodel.Upsert
// @router /categories/{id} [PUT]
func (categoryImpl) Update(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		staff   = echocontext.GetStaff(c).(*auth.User)
		s       = service.Category(staff)
		id      = echocontext.GetParam(c, "id").(string)
		payload = echocontext.GetPayload(c).(requestmodel.CategoryUpdatePayload)
	)

	res, err := s.Update(ctx, id, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}
