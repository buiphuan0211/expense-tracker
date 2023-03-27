package dao

import (
	"context"
	mgmodel "expense-tracker/internal/model/mg"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StaffInterface interface {
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgmodel.Staff)
	FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgmodel.Staff, err error)
	InsertOne(ctx context.Context, payload mgmodel.Staff) (err error)
	CountByCondition(ctx context.Context, cond interface{}) (total int64)
}

type staffImpl struct{}

func Staff() StaffInterface {
	return staffImpl{}
}
