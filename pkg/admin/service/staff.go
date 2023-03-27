package service

import (
	"context"
	"expense-tracker/internal/auth"
	mgmodel "expense-tracker/internal/model/mg"
	"expense-tracker/internal/util/mgquery"
	responsemodel "expense-tracker/pkg/admin/model/response"
)

// StaffInterface ...
type StaffInterface interface {
	All(ctx context.Context, q mgquery.AppQuery) (res responsemodel.StaffAll)
	ExistedByEmail(ctx context.Context, username string) (res bool)
	FindByEmail(ctx context.Context, email string) (res mgmodel.Staff, err error)
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
