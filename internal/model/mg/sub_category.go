package mgmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// SubCategory ...
type SubCategory struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Slug         string             `bson:"slug"`
	SearchString string             `bson:"searchString"`
	Category     primitive.ObjectID `bson:"category"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	CreatedBy    primitive.ObjectID `bson:"createdBy"`
}
