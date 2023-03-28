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

// UpdatePermissions ...
func (s staffImpl) UpdatePermissions(ctx context.Context, id string, payload requestmodel.StaffPermissionUpdatePayload) (res responsemodel.Upsert, err error) {
	if existed := s.ExistedByID(ctx, id); !existed {
		err = errors.New(errorcode.StaffNotFound)
		return
	}

	var (
		d      = dao.Staff()
		update = bson.M{"$set": payload.ConvertToBSON()}
		cond   = bson.M{"_id": pconvert.StringToObjectID(id)}
	)

	err = d.UpdateOneByCondition(ctx, cond, update)
	if err != nil {
		err = errors.New(errorcode.StaffErrorWhenUpdatePermissions)
		return
	}

	res.ID = id
	return
}
