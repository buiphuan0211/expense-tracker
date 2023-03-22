package requestmodel

import (
	mgmodel "expense-tracker/internal/model/mg"
	"expense-tracker/internal/util/pmongo"
	"expense-tracker/internal/util/pstring"
	"expense-tracker/internal/util/ptime"
	"expense-tracker/pkg/admin/errorcode"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
)

//
// Create
//

// CategoryCreatePayload ...
type CategoryCreatePayload struct {
	Name string `json:"name"`
}

func (m CategoryCreatePayload) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required.Error(errorcode.CategoryRequiredName)),
	)
}

// ConvertToBSON ...
func (m CategoryCreatePayload) ConvertToBSON() mgmodel.Category {
	return mgmodel.Category{
		ID:           pmongo.NewObjectID(),
		Name:         m.Name,
		Slug:         pstring.ToSlug(strings.ToLower(m.Name)),
		SearchString: pstring.NonAccentVietnamese(m.Name),
		CreatedAt:    ptime.Now(),
		UpdatedAt:    ptime.Now(),
	}
}

//
// All
//

// CategoryAll ...
type CategoryAll struct {
	Page    int64  `query:"page"`
	Limit   int64  `query:"limit"`
	Keyword string `query:"keyword"`
}

func (m CategoryAll) Validate() error {
	return validation.ValidateStruct(&m)
}
