package dao

import (
	"context"
	"expense-tracker/internal/config/database"
	"expense-tracker/internal/config/plogger"
)

// UpdateOneByCondition ...
func (categoryImpl) UpdateOneByCondition(ctx context.Context, cond interface{}, payload interface{}) (err error) {
	var col = database.GetCategoryCol()
	if _, err = col.UpdateByID(ctx, cond, payload); err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.Category - UpdateByCondition",
			Message: err.Error(),
			Data: plogger.Map{
				"cond":    cond,
				"payload": payload,
			},
		})
	}

	return
}
