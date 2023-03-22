package dao

import (
	"context"
	"expense-tracker/internal/config/database"
	"expense-tracker/internal/config/plogger"
	mgmodel "expense-tracker/internal/model/mg"
)

// Create ...
func (categoryImpl) Create(ctx context.Context, payload mgmodel.Category) error {
	var col = database.GetCategoryCol()
	if _, err := col.InsertOne(ctx, payload); err != nil {
		plogger.Error("dao.Category - InsertOne", plogger.LogData{
			Source:  "expense-tracker",
			Message: err.Error(),
			Data: plogger.Map{
				"payload": payload,
			},
		})
	}

	return nil
}
