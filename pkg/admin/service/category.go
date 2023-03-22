package service

import (
	"context"
	"expense-tracker/internal/util/mgquery"
	requestmodel "expense-tracker/pkg/admin/model/request"
	responsemodel "expense-tracker/pkg/admin/model/response"
)

// CategoryInterface ...
type CategoryInterface interface {
	Create(ctx context.Context, payload requestmodel.CategoryCreatePayload) (res responsemodel.Upsert, err error)
	All(ctx context.Context, q mgquery.AppQuery) (res responsemodel.CategoryAll)
}

// categoryImpl ...
type categoryImpl struct{}

// Category ...
func Category() CategoryInterface {
	return categoryImpl{}
}
