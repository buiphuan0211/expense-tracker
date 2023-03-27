package requestmodel

import (
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
	return validation.ValidateStruct(&m)
}
