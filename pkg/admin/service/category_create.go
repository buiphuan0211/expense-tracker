package service

import (
	"context"
	"errors"
	"expense-tracker/pkg/admin/dao"
	"expense-tracker/pkg/admin/errorcode"
	requestmodel "expense-tracker/pkg/admin/model/request"
	responsemodel "expense-tracker/pkg/admin/model/response"
)

// Create ...
func (s categoryImpl) Create(ctx context.Context, payload requestmodel.CategoryCreatePayload) (res responsemodel.Upsert, err error) {
	// Check existed categoryName
	if isTrue := s.IsExistedByName(ctx, payload.Name); isTrue {
		err = errors.New(errorcode.CategoryExistedName)
		return
	}

	var (
		d   = dao.Category()
		doc = payload.ConvertToBSON()
	)

	if err = d.Create(ctx, doc); err != nil {
		return
	}

	res.ID = doc.ID.Hex()
	return
}
