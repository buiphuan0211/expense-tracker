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
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	Account      StaffAccount       `bson:"account"`
}

// StaffAccount ...
type StaffAccount struct {
	Username   string   `bson:"username"`
	Password   string   `bson:"password"`
	Permission []string `bson:"permission"`
	IsRoot     bool     `bson:"isRoot"`
}
