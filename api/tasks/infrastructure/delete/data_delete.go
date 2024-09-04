package tasks_delete_infrastructure

import (
	tasks_domain "go-api/api/tasks/domain"
	tasks_delete_db_infrastructure "go-api/api/tasks/infrastructure/delete/db"
	tasks_response_infrastructure "go-api/api/tasks/infrastructure/response"
	"go-api/logs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// TaskHandler is a structure that handles operations related to tasks.
type TaskHandler struct {
	db interface{} // The 'db' field can be of any type since it's an empty interface
}

// NewTaskHander creates a new instance of TaskHandler with the provided database object.
// Returns a type that implements the tasks_domain.DeleteTaskInterface interface.
func NewTaskHander(db interface{}) tasks_domain.DeleteTaskInterface {
	// Create a new instance of TaskHandler with the provided database object
	return &TaskHandler{db: db}
}

// DeleteData handles the request to delete a Task.
// @Summary Delete a task by ID.
// @Description Delete a task from the system using its ID.
// @ID delete-task
// @Produce json
// @Param id path string true "ID of the task to delete"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal server error"
// @Router /tasks/{id} [delete]
func (th *TaskHandler) DeleteData(c *gin.Context) {
	// Get the ID from the parameter in the URL
	id := c.Param("id")

	var post tasks_domain.Task

	// Try to bind the JSON from the request to the Task object
	if err := c.BindJSON(&post); err != nil {
		// If there's a JSON format error, respond with a bad request error
		tasks_response_infrastructure.CreateBadResponse(c)
		logs.Warning_Logger.Println("JSON format error")
		return // End the function execution
	}

	// Select the type of database based on the type of `db` in `TaskHandler`
	switch db := th.db.(type) {
	case *gorm.DB:
		// If `db` is of type *gorm.DB (SQL database), call the SQL delete function
		delete := tasks_delete_db_infrastructure.DeleteTaskSql(db, id)
		if !delete {
			// If deletion fails, respond with a delete error
			tasks_response_infrastructure.CreateDeleteErrorResponse(c)
			return // End the function execution
		}
		// If deletion is successful, respond with a delete confirmation
		tasks_response_infrastructure.CreateDeleteResponse(c)
		break
	case *mongo.Database:
		// If `db` is of type *mongo.Database (MongoDB database), log a warning message
		logs.Warning_Logger.Println("No support for deletion in MongoDB")
		break
	default:
		// If `db` is not of a compatible type, log a warning message
		logs.Warning_Logger.Println("Unsupported database type")
		break
	}
}
