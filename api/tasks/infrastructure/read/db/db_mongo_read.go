package tasks_read_db_infrastructure

import (
	"context"
	tasks_domain "go-api/api/tasks/domain"
	"go-api/logs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllTaskMongo gets all tasks from the "tasks" collection in MongoDB.
func GetAllTaskMongo(mongodb *mongo.Database) ([]tasks_domain.Task, error) {
	// Variable to store retrieved tasks
	var tasks []tasks_domain.Task

	// Perform a query to get all tasks from the "tasks" collection
	cur, err := mongodb.Collection("tasks").Find(context.Background(), bson.M{})
	if err != nil {
		// If there's an error during the query, log the error in the error logger and return the error
		logs.Error_Logger.Fatal(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	// Iterate through the retrieved results
	for cur.Next(context.Background()) {
		// Variable to store an individual task
		var task tasks_domain.Task
		// Decode the result into the Task structure
		err := cur.Decode(&task)
		if err != nil {
			// If there's an error during decoding, log the error in the error logger and return the error
			logs.Error_Logger.Fatal(err)
			return nil, err
		}
		// Add the task to the tasks list
		tasks = append(tasks, task)
	}

	// Check if there was any error during the traversal of results
	if err := cur.Err(); err != nil {
		// If there's an error, log the error in the error logger and return the error
		logs.Error_Logger.Fatal(err)
		return nil, err
	}

	// Log the cursor results
	logs.Info_Logger.Println(cur)

	// Return the tasks and nil, indicating no error occurred
	return tasks, nil
}
