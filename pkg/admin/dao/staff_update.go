package dao

import (
	"context"
	"expense-tracker/internal/config/database"
	"expense-tracker/internal/config/plogger"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateOneByCondition ...
func (staffImpl) UpdateOneByCondition(ctx context.Context, cond interface{}, payload interface{}, opts ...*options.UpdateOptions) (err error) {
	var col = database.GetStaffCol()
	if _, err = col.UpdateOne(ctx, cond, payload, opts...); err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.Staff - UpdateOneByCondition",
			Message: err.Error(),
			Data: plogger.Map{
				"cond":    cond,
				"payload": payload,
			},
		})
	}
	return
}
