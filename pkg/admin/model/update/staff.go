package updatemodel

import "time"

// StaffPermissionsUpdate ...
type StaffPermissionsUpdate struct {
	Permissions []string  `bson:"permissions"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}
