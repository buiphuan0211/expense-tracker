package responsemodel

import (
	"expense-tracker/internal/util/ptime"
)

// CategoryAll ...
type CategoryAll struct {
	List  []CategoryBrief `json:"list"`
	Limit int64           `json:"limit"`
	Total int64           `json:"total"`
}

// CategoryBrief ..
type CategoryBrief struct {
	ID        string              `json:"_id"`
	Name      string              `json:"name"`
	CreatedAt *ptime.TimeResponse `json:"createdAt"`
}

// CategoryDetail ...
type CategoryDetail struct {
	ID        string              `json:"_id"`
	Name      string              `json:"name"`
	CreatedAt *ptime.TimeResponse `json:"createdAt"`
}
