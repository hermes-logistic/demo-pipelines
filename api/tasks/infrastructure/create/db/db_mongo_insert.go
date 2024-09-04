package tasks_create_db_infrastructure

import (
	"context"
	"go-api/logs"

	"go.mongodb.org/mongo-driver/mongo"
)

// InsertTaskMongo inserts a new document into the "tasks" collection in MongoDB.
func InsertTaskMongo(mongodb *mongo.Database, data NoSqlTask) bool {
	// Create a background context
	ctx := context.Background()

	// Insert the 'data' document into the "tasks" collection
	insert, err := mongodb.Collection("tasks").InsertOne(ctx, data)

	if err != nil {
		// If there's an error during insertion, log the error in the error logger
		logs.Error_Logger.Println(err)
		return false // Return false indicating insertion failed
	} else {
		// If insertion is successful, log the ID of the inserted document in the info logger
		logs.Info_Logger.Printf("Data insertion completed: %v\n", insert.InsertedID)
		return true // Return true indicating insertion was successful
	}
}
