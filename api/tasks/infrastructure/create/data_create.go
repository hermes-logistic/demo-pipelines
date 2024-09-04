package tasks_create_infrastructure

import (
	"fmt"
	tasks_empty_application "go-api/api/tasks/application/empty"
	tasks_domain "go-api/api/tasks/domain"
	tasks_create_db_infrastructure "go-api/api/tasks/infrastructure/create/db"
	tasks_response_infrastructure "go-api/api/tasks/infrastructure/response"
	"go-api/logs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// TaskHandler handles operations related to tasks.
type TaskHandler struct {
	db interface{} // The empty interface allows TaskHandler to handle different types of databases
}

// NewTaskData creates a new instance of TaskHandler with the provided database.
func NewTaskData(db interface{}) tasks_domain.CreateTaskInterface {
	return &TaskHandler{db: db}
}

// CreateData handles the creation of a new task.
// @Summary Create a new task
// @Description Create a new task with the data provided in the JSON body
// @Accept json
// @Produce json
// @Param task body tasks_domain.Task true "Data of the new task"
// @Success 200 {object} tasks_domain.Task "Inserted task"
// @Failure 400 {object} tasks_domain.Task "Error in input data"
// @Failure 500 {object} tasks_domain.Task "Error in input data"
// @Router /tasks [post]
func (th TaskHandler) CreateData(c *gin.Context) {
	// Generate a new namespace and ID for the task
	namespace := uuid.New()
	logs.Info_Logger.Println(namespace.String())
	name := []byte("tasks")
	id := uuid.NewSHA1(namespace, name)

	// Create a structure to store the task data received in the JSON body
	var post tasks_domain.Task

	// Try to bind the received JSON data to the Task struct
	if err := c.BindJSON(&post); err != nil {
		// If there is a JSON format error, respond with an error message
		tasks_response_infrastructure.CreateBadResponse(c)
		logs.Warning_Logger.Println("Format error")
		return
	}

	// Validate that the task does not have empty fields
	validate := tasks_empty_application.EmptyValidation(&post)
	if !validate {
		// If validation fails, respond with an error message
		tasks_response_infrastructure.CreateErrorResponse(c)
		logs.Warning_Logger.Println("Empty fields")
		return
	}

	// Get the type of the provided database for the handler
	typeToEvaluate := fmt.Sprintf("%T", th.db)
	fmt.Println("Type to evaluate:", typeToEvaluate)

	// Perform a switch based on the type of the database
	switch db := th.db.(type) {
	case *gorm.DB:
		// If the database is of type *gorm.DB, create a SqlTask object to insert into the SQL database
		insertData := tasks_create_db_infrastructure.SqlTask{
			ID:     id.String(),
			Name:   post.Name,
			Status: post.Status,
		}

		// Insert the task into the SQL database using the InsertTask function
		insert := tasks_create_db_infrastructure.InsertTask(*db, insertData)

		// Check if the insertion was successful
		if !insert {
			// If the insertion fails, respond with an error message
			tasks_response_infrastructure.CreateErrorResponse(c)
			return
		}

		// If the insertion was successful, respond with a confirmation message
		tasks_response_infrastructure.CreateConfirmResponse(c)
	case *mongo.Database:
		// If the database is of type *mongo.Database, create a NoSqlTask object to insert into the NoSQL database
		insertMongo := tasks_create_db_infrastructure.NoSqlTask{
			ID:     id.String(),
			Name:   post.Name,
			Status: post.Status,
		}

		// Insert the task into the NoSQL database using the InsertTaskMongo function
		mongoInsert := tasks_create_db_infrastructure.InsertTaskMongo(db, insertMongo)

		// Check if the insertion was successful
		if !mongoInsert {
			// If the insertion fails, respond with an error message
			tasks_response_infrastructure.CreateErrorResponse(c)
			return
		}

		// If the insertion was successful, respond with a confirmation message
		tasks_response_infrastructure.CreateConfirmResponse(c)
	default:
		// If the database type is not supported, log an information message
		logs.Info_Logger.Println("Unsupported database")
	}
}
