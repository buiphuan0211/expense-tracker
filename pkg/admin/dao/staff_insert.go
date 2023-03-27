package dao

import (
	"context"
	"expense-tracker/internal/config/database"
	"expense-tracker/internal/config/plogger"
	mgmodel "expense-tracker/internal/model/mg"
)

// InsertOne ...
func (staffImpl) InsertOne(ctx context.Context, payload mgmodel.Staff) (err error) {
	var col = database.GetStaffCol()
	if _, err = col.InsertOne(ctx, payload); err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "dao.Staff - InserOne",
			Message: err.Error(),
			Data: plogger.Map{
				"payload": payload,
			},
		})
	}
	return
}
