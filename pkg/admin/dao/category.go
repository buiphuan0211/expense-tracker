package dao

import (
	"context"
	mgmodel "expense-tracker/internal/model/mg"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CategoryInterface ...
type CategoryInterface interface {
	Create(ctx context.Context, payload mgmodel.Category) error
	FindOnByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgmodel.Category, err error)
	FindByID(ctx context.Context, id primitive.ObjectID) (doc *mgmodel.Category, err error)
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgmodel.Category)
	CountByCondition(ctx context.Context, cond interface{}) (total int64)
	UpdateOneByCondition(ctx context.Context, cond interface{}, payload interface{}) (err error)
}

// categoryImpl ...
type categoryImpl struct{}

// Category ...
func Category() CategoryInterface {
	return categoryImpl{}
}
