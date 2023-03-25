package updatemodel

import "time"

// CategoryUpdate ...
type CategoryUpdate struct {
	Name         string    `bson:"name"`
	Slug         string    `bson:"slug"`
	SearchString string    `bson:"searchString"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
