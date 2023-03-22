package dao

import (
	"context"
	"expense-tracker/internal/config/database"
	"expense-tracker/internal/config/plogger"
	mgmodel "expense-tracker/internal/model/mg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOnByCondition ...
func (categoryImpl) FindOnByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgmodel.Category, err error) {
	var col = database.GetCategoryCol()
	if err = col.FindOne(ctx, cond, opts...).Decode(&doc); err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.Category - FindOnByCondition",
			Message: err.Error(),
			Data: plogger.Map{
				"cond": cond,
			},
		})
	}

	return
}

// FindByID ...
func (categoryImpl) FindByID(ctx context.Context, id primitive.ObjectID) (doc *mgmodel.Category, err error) {
	var col = database.GetCategoryCol()
	if err = col.FindOne(ctx, bson.M{"_id": id}).Decode(&doc); err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.Category - FindByID",
			Message: err.Error(),
			Data: plogger.Map{
				"id": id,
			},
		})
	}

	return
}

// FindByCondition ...
func (categoryImpl) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgmodel.Category) {
	var col = database.GetCategoryCol()
	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.category - FindByCondition",
			Message: err.Error(),
			Data: plogger.Map{
				"cond": cond,
				"opts": opts,
			},
		})
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &docs); err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.category - FindByCondition",
			Message: err.Error(),
			Data: plogger.Map{
				"cond": cond,
				"opts": opts,
			},
		})
	}

	return
}

// CountByCondition ...
func (categoryImpl) CountByCondition(ctx context.Context, cond interface{}) (total int64) {
	var col = database.GetCategoryCol()
	total, _ = col.CountDocuments(ctx, cond)
	return
}
