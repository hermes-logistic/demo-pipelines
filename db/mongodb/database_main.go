package mongo_db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Data structure that holds the context of the database
type Mongo struct {
	DB *mongo.Database
}
