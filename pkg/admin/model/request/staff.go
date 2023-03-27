package requestmodel

import (
	"expense-tracker/internal/constant"
	"expense-tracker/pkg/admin/errorcode"
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
		validation.Field(&m.Status, validation.Each(validation.In(statuses...).Error(errorcode.StaffInvalidSStatus))))
}
