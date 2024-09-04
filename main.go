package main

import (
	mongo_db "go-api/db/mongodb"
	sql_db "go-api/db/sql"
	_ "go-api/docs"
	"go-api/logs"
	"go-api/router"
	"os"
)

// @title backend api with go
func main() {
	// Get the context from the environment
	context := os.Getenv("context")

	// Select action based on the context
	switch context {
	case "mongo":
		// Initialize a Mongo instance and establish the connection
		var d mongo_db.Mongo
		d.MongoConnection()

		// Initialize a Routes instance
		var r router.Routes

		// Configure the routes and run the server
		r.CreateRoutes(d.DB)
		r.Run()
	case "postgres":
		// Initialize a Database instance for PostgreSQL and connect
		var d sql_db.Database
		d.Connect(os.Getenv("DB_STRING"))

		// Run database migrations
		d.Migration()

		// Initialize a Routes instance
		var r router.Routes

		// Configure the routes and run the server
		r.CreateRoutes(d.DB)
		r.Run()
	case "local":
		var local router.Routes

		local.CreateRoutes("null")
		local.Run()
	default:
		// Handle an invalid context
		logs.Error_Logger.Fatal("Invalid Context")
	}
}
