package service

import (
	"context"
	"expense-tracker/internal/config/plogger"
	mgmodel "expense-tracker/internal/model/mg"
	pgenerate "expense-tracker/internal/util/generate"
	"expense-tracker/internal/util/ptime"
	"expense-tracker/pkg/admin/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MigrateInterface ...
type MigrateInterface interface {
	MigrationStaffs()
}

// MigrateImpl ...
type MigrateImpl struct {
}

// Migrate ...
func Migrate() MigrateInterface {
	return MigrateImpl{}
}

// MigrationStaffs ...
func (MigrateImpl) MigrationStaffs() {
	var (
		d   = dao.Staff()
		ctx = context.Background()
	)
	payload := mgmodel.Staff{
		ID:           primitive.NewObjectID(),
		Name:         "Mèo",
		SearchString: "meo",
		Phone:        "0384062418",
		Email:        "meo@gmail.com",
		Password:     pgenerate.HashPassword("1234"),
		Permission:   []string{"staff_view", "staff_edit"},
		IsRoot:       true,
		CreatedAt:    ptime.Now(),
		UpdatedAt:    ptime.Now(),
	}

	if err := d.InsertOne(ctx, payload); err != nil {
		plogger.Error("", plogger.LogData{
			Source:  "migration - MigrationStaffs",
			Message: err.Error(),
			Data: plogger.Map{
				"payload": payload,
			},
		})
	}

}
