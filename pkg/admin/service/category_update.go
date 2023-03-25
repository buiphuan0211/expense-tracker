package service

import (
	"context"
	"errors"
	"expense-tracker/internal/util/pconvert"
	"expense-tracker/pkg/admin/dao"
	"expense-tracker/pkg/admin/errorcode"
	requestmodel "expense-tracker/pkg/admin/model/request"
	responsemodel "expense-tracker/pkg/admin/model/response"
	"go.mongodb.org/mongo-driver/bson"
)

// Update ...
func (s categoryImpl) Update(ctx context.Context, id string, payload requestmodel.CategoryUpdatePayload) (res responsemodel.Upsert, err error) {
	if isExisted := s.IsExistedByID(ctx, id); !isExisted {
		err = errors.New(errorcode.CategoryNotFound)
		return
	}

	var (
		d      = dao.Category()
		objID  = pconvert.StringToObjectID(id)
		cond   = bson.M{"_id": objID}
		update = bson.M{"$set": payload.ConvertToBSON()}
	)

	if err = d.UpdateOneByCondition(ctx, cond, update); err != nil {
		return
	}

	res.ID = id
	return
}
