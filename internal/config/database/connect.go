package database

import (
	"context"
	"expense-tracker/internal/config"
	"fmt"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

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
	db = client.Database(cfg.DBName)
	index()
}
