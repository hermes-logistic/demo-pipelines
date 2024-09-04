package tasks_read_infrastructure

import (
	tasks_domain "go-api/api/tasks/domain"
	tasks_read_db_infrastructure "go-api/api/tasks/infrastructure/read/db"
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

// NewTaskHandler creates a new instance of TaskHandler with the provided database object.
// Returns a type that implements the tasks_domain.ReadAllTaskInterface interface.
func NewTaskHandler(db interface{}) tasks_domain.ReadAllTaskInterface {
	// Create a new instance of TaskHandler with the provided database object
	return &TaskHandler{db: db}
}

// ReadData handles the request to read all tasks.
// @Summary Get all tasks.
// @Description Get all tasks from the system.
// @ID get-all-tasks
// @Produce json
// @Success 200 {array} tasks_domain.Task "OK"
// @Failure 500 {string} string "Internal server error"
// @Router /tasks [get]
func (th *TaskHandler) ReadData(c *gin.Context) {
	// Select the type of database based on the type of `db` in `TaskHandler`
	switch db := th.db.(type) {
	case *gorm.DB:
		// If `db` is of type *gorm.DB (SQL database), get all tasks from the SQL database
		tasks, err := tasks_read_db_infrastructure.GetAllTaskSql(db)
		if err != nil {
			// If there's an error getting tasks, respond with a read error
			tasks_response_infrastructure.CreateReadErrorResponse(c)
			return
		}
		// If tasks are retrieved successfully, respond with the tasks
		tasks_response_infrastructure.CreateReadResponse(c, tasks)
	case *mongo.Database:
		// If `db` is of type *mongo.Database (MongoDB database), get all tasks from the MongoDB database
		tasks, err := tasks_read_db_infrastructure.GetAllTaskMongo(db)
		if err != nil {
			// If there's an error getting tasks, respond with a read error
			tasks_response_infrastructure.CreateReadErrorResponse(c)
			return
		}
		// If tasks are retrieved successfully, respond with the tasks
		tasks_response_infrastructure.CreateReadResponse(c, tasks)
	default:
		// If `db` is not of a compatible type, log a warning message
		logs.Info_Logger.Println("No support for this database")
	}
}
