package service

import (
	"context"
	"expense-tracker/internal/auth"
	"expense-tracker/internal/util/mgquery"
	requestmodel "expense-tracker/pkg/admin/model/request"
	responsemodel "expense-tracker/pkg/admin/model/response"
)

// CategoryInterface ...
type CategoryInterface interface {
	Create(ctx context.Context, payload requestmodel.CategoryCreatePayload) (res responsemodel.Upsert, err error)
	All(ctx context.Context, q mgquery.AppQuery) (res responsemodel.CategoryAll)
	Detail(ctx context.Context, id string) (res responsemodel.CategoryDetail, err error)
	Update(ctx context.Context, id string, payload requestmodel.CategoryUpdatePayload) (res responsemodel.Upsert, err error)
}

// categoryImpl ...
type categoryImpl struct {
	staff *auth.User
}

// Category ...
func Category(cs *auth.User) CategoryInterface {
	return categoryImpl{
		staff: cs,
	}
}
