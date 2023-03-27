package mgmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Staff ...
type Staff struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	SearchString string             `bson:"searchString"`
	Phone        string             `bson:"phone"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	Permission   []string           `bson:"permission"`
	IsRoot       bool               `bson:"isRoot"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}
