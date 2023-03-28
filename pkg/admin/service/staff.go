package service

import (
	"context"
	"expense-tracker/internal/auth"
	mgmodel "expense-tracker/internal/model/mg"
	"expense-tracker/internal/util/mgquery"
	requestmodel "expense-tracker/pkg/admin/model/request"
	responsemodel "expense-tracker/pkg/admin/model/response"
)

// StaffInterface ...
type StaffInterface interface {
	All(ctx context.Context, q mgquery.AppQuery) (res responsemodel.StaffAll)
	ExistedByEmail(ctx context.Context, username string) (res bool)
	FindRawByEmail(ctx context.Context, email string) (res mgmodel.Staff, err error)
	ExistedByID(ctx context.Context, id string) (res bool)
	UpdatePermissions(ctx context.Context, id string, payload requestmodel.StaffPermissionUpdatePayload) (res responsemodel.Upsert, err error)
}

// staffImpl ...
type staffImpl struct {
	staff *auth.User
}

// Staff ...
func Staff(cs *auth.User) StaffInterface {
	return staffImpl{
		staff: cs,
	}
}
