package database

import (
	"context"
	"expense-tracker/internal/config"
	"fmt"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var expenseTrackerDB *mongo.Database

// ConnectMongoDBExpenseTracker ...
func ConnectMongoDBExpenseTracker() {
	cfg := config.GetENV().MongoDB
	connectOptions := options.ClientOptions{}
	client, err := mongo.Connect(context.Background(), connectOptions.ApplyURI(cfg.URI))
	if err != nil {
		fmt.Println("Error when connect to MongoDB database", cfg.URI, err)
		panic(err)
	}

	fmt.Println(aurora.Green("*** CONNECTED TO MONGODB: " + cfg.URI + " --- DB: " + cfg.DBName))
	expenseTrackerDB = client.Database(cfg.DBName)
	index()
}

func GetMongoDBExpenseTracker() *mongo.Database {
	return expenseTrackerDB
}

func GetCategoryCol() *mongo.Collection {
	return expenseTrackerDB.Collection(colCategory)
}
