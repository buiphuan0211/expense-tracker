package dao

import (
	"context"
	"expense-tracker/internal/config/database"
	"expense-tracker/internal/config/plogger"
	mgmodel "expense-tracker/internal/model/mg"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindByCondition ...
func (staffImpl) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgmodel.Staff) {
	var col = database.GetStaffCol()
	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.Staff - FindByCondition - Find",
			Message: err.Error(),
			Data: plogger.Map{
				"cond": cond,
				"opts": opts,
			},
		})
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &docs); err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.Staff - FindByCondition - All",
			Message: err.Error(),
			Data: plogger.Map{
				"cond": cond,
				"opts": opts,
			},
		})
	}

	return
}

// FindOneByCondition ...
func (staffImpl) FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgmodel.Staff, err error) {
	var col = database.GetStaffCol()
	if err = col.FindOne(ctx, cond, opts...).Decode(&doc); err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.Staff - FindOneByCondition",
			Message: err.Error(),
			Data: plogger.Map{
				"cond": cond,
			},
		})
	}

	return
}

// CountByCondition ...
func (staffImpl) CountByCondition(ctx context.Context, cond interface{}) (total int64) {
	var col = database.GetStaffCol()
	total, _ = col.CountDocuments(ctx, cond)
	return
}
