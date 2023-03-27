package service

import (
	"context"
	"errors"
	pgenerate "expense-tracker/internal/util/generate"
	"expense-tracker/pkg/admin/dao"
	"expense-tracker/pkg/admin/errorcode"
	requestmodel "expense-tracker/pkg/admin/model/request"
)

// StaffAuthInterface ...
type StaffAuthInterface interface {
	Register(ctx context.Context, payload requestmodel.RegisterPayload) (token string, err error)
	Login(ctx context.Context, payload requestmodel.LoginPayload) (token string, err error)
}

// StaffAuthImpl ...
type StaffAuthImpl struct{}

// StaffAuth ...
func StaffAuth() StaffAuthInterface {
	return StaffAuthImpl{}
}

// Register ...
func (s StaffAuthImpl) Register(ctx context.Context, payload requestmodel.RegisterPayload) (token string, err error) {
	var staffSvc = Staff(nil)

	// Check existed by email
	if existed := staffSvc.ExistedByEmail(ctx, payload.Email); existed {
		err = errors.New(errorcode.StaffExistedEmail)
		return
	}

	// Insert
	var (
		d   = dao.Staff()
		doc = payload.ConvertToBSON()
	)
	if err = d.InsertOne(ctx, doc); err != nil {
		return
	}

	// Generate token
	var payloadToken = requestmodel.TokenPayload{
		ID:     doc.ID.Hex(),
		Name:   doc.Name,
		IsRoot: false,
	}

	if token, err = pgenerate.TokenString(payloadToken); err != nil {
		err = errors.New(errorcode.StaffAuthGenerateToken)
		return
	}

	return
}

// Login ...
func (s StaffAuthImpl) Login(ctx context.Context, payload requestmodel.LoginPayload) (token string, err error) {
	var staffSvc = Staff(nil)

	// Check invalid email and password
	if existed := staffSvc.ExistedByEmail(ctx, payload.Email); !existed {
		err = errors.New(errorcode.StaffAuthInvalidEmailOrPassword)
		return
	}

	staff, err := staffSvc.FindRawByEmail(ctx, payload.Email)
	if err != nil {
		err = errors.New(errorcode.StaffNotFound)
		return
	}

	if isValid := pgenerate.CheckPasswordHash(staff.Password, payload.Password); !isValid {
		err = errors.New(errorcode.StaffAuthInvalidEmailOrPassword)
		return
	}

	// Generate token
	var payloadToken = requestmodel.TokenPayload{
		ID:          staff.ID.Hex(),
		Name:        staff.Name,
		Phone:       staff.Phone,
		Email:       staff.Email,
		Permissions: staff.Permission,
		IsRoot:      staff.IsRoot,
	}

	if token, err = pgenerate.TokenString(payloadToken); err != nil {
		err = errors.New(errorcode.StaffAuthGenerateToken)
		return
	}

	return
}
