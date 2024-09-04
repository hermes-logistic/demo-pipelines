package mongo_db

import (
	"context"
	"go-api/logs"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection establishes a connection to the MongoDB database.
func (db *Mongo) MongoConnection() {
	// Get the MongoDB connection URI from the "MONGO_URI" environment variable
	uri := os.Getenv("MONGO_URI")

	// Create a background context
	ctx := context.Background()

	// Connect to the MongoDB server using the URI and get a client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		// In case of error connecting, log the error and exit the application
		logs.Error_Logger.Fatal(err)
	}

	// Check the connection to the server
	err = client.Ping(ctx, nil)
	if err != nil {
		// In case of error pinging, log the error and exit the application
		logs.Error_Logger.Fatal(err)
	}

	// If the connection was successful, log a success message
	logs.Info_Logger.Println("Successful connection")

	// Select the "go-api" database in MongoDB
	database := client.Database("go-api")

	// Assign the database connection to the DB field of the Mongo structure
	db.DB = database
}
