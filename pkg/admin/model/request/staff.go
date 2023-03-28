package requestmodel

import (
	"expense-tracker/internal/constant"
	"expense-tracker/internal/util/ptime"
	"expense-tracker/pkg/admin/errorcode"
	updatemodel "expense-tracker/pkg/admin/model/update"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// StaffAll ...
type StaffAll struct {
	Page    int64  `query:"page"`
	Limit   int64  `query:"limit"`
	Keyword string `query:"keyword"`
	Status  string `query:"status"`
}

func (m StaffAll) Validate() error {
	statuses := []interface{}{
		constant.StatusActive,
		constant.StatusInActive,
	}
	return validation.ValidateStruct(&m,
		validation.Field(&m.Status, validation.In(statuses...).Error(errorcode.StaffInvalidStatus)))
}

// StaffPermissionUpdatePayload ...
type StaffPermissionUpdatePayload struct {
	Permissions []string `query:"permissions"`
}

func (m StaffPermissionUpdatePayload) Validate() error {
	return validation.ValidateStruct(&m)
}

// ConvertToBSON ...
func (m StaffPermissionUpdatePayload) ConvertToBSON() updatemodel.StaffPermissionsUpdate {
	return updatemodel.StaffPermissionsUpdate{
		Permissions: m.Permissions,
		UpdatedAt:   ptime.Now(),
	}
}
