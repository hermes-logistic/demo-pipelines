package tasks_update_infrastructure

import (
	tasks_domain "go-api/api/tasks/domain"
	tasks_response_infrastructure "go-api/api/tasks/infrastructure/response"
	tasks_update_db_infrastructure "go-api/api/tasks/infrastructure/update/db"
	"go-api/logs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// TaskHandler is a structure that handles operations related to tasks.
type TaskHandler struct {
	db interface{} // The db field can be of any type since it's an empty interface
}

// NewTaskHandler creates a new instance of TaskHandler with the provided database object.
// Returns a type that implements the tasks_domain.UpdateTaskInterface interface.
func NewTaskHandler(db interface{}) tasks_domain.UpdateTaskInterface {
	// Create a new instance of TaskHandler with the provided database object
	return &TaskHandler{db: db}
}

// UpdateData handles the request to update a task.
// @Summary Update a task by ID.
// @Description Update a task in the system using its ID and the new provided data.
// @ID update-task
// @Accept json
// @Produce json
// @Param id path string true "ID of the task to update"
// @Param task body tasks_domain.Task true "New data of the task"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/{id} [put]
func (th *TaskHandler) UpdateData(c *gin.Context) {
	// Get the ID of the task to update from the URL parameters
	id := c.Param("id")

	// Create a new task based on the data received in the JSON request
	var post tasks_domain.Task
	if err := c.BindJSON(&post); err != nil {
		// If there's a JSON format error, respond with a bad request error
		tasks_response_infrastructure.CreateBadResponse(c)
		logs.Warning_Logger.Println("JSON format error")
		return
	}

	// Select the database type based on the type of `db` in `TaskHandler`
	switch db := th.db.(type) {
	case *gorm.DB:
		// If `db` is of type *gorm.DB (SQL database), create a Task object with the updated data
		data := tasks_domain.Task{
			ID:     id,
			Name:   post.Name,
			Status: post.Status,
		}

		// Call the SQL update function
		update := tasks_update_db_infrastructure.UpdateTaskSql(*db, data)

		if !update {
			// If the update fails, respond with an update error
			tasks_response_infrastructure.CreateUpdateErrorResponse(c)
			return
		}

		// If the update is successful, respond with an update confirmation
		tasks_response_infrastructure.CreateUpdateResponse(c)
		break
	case *mongo.Database:
		// If `db` is of type *mongo.Database (MongoDB), log a warning message
		logs.Info_Logger.Println("Unsupported database")
		break
	default:
		// If `db` is not of a supported type, log a warning message
		logs.Info_Logger.Println("Unsupported database")
		break
	}
}
