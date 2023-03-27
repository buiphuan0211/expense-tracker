package responsemodel

import (
	"expense-tracker/internal/util/ptime"
)

// StaffAll ...
type StaffAll struct {
	List  []StaffBrief `json:"list"`
	Total int64        `json:"page"`
	Limit int64        `json:"limit"`
}

// StaffBrief ...
type StaffBrief struct {
	ID         string              `json:"_id"`
	Name       string              `json:"name"`
	Phone      string              `json:"phone"`
	Email      string              `json:"email"`
	Permission []string            `json:"permission"`
	IsRoot     bool                `json:"isRoot"`
	CreatedAt  *ptime.TimeResponse `json:"createdAt"`
}
