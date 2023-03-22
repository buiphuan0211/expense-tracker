package handler

import (
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
		s       = service.Category()
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
		s       = service.Category()
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
