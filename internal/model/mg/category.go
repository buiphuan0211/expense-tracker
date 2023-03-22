package mgmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Category ...
type Category struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Slug         string             `bson:"slug"`
	SearchString string             `bson:"searchString"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	CreatedBy    primitive.ObjectID `bson:"createdBy"`
}
