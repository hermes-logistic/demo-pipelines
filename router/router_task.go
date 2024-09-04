package router

import (
	tasks_create_infrastructure "go-api/api/tasks/infrastructure/create"
	tasks_read_infrastructure "go-api/api/tasks/infrastructure/read"

	"github.com/gin-gonic/gin"
)

// TasksRoutes configures routes related to tasks on the Gin router.
func (ro *Routes) TasksRoutes(r *gin.Engine, db interface{}) {
	// Create a new task handler to read task data
	tr := tasks_read_infrastructure.NewTaskHandler(db)
	// Configure the GET route to get all tasks
	r.GET("/tasks", tr.ReadData)

	// Create a new task handler to create task data
	tc := tasks_create_infrastructure.NewTaskData(db)
	// Configure the POST route to create a new task
	r.POST("/tasks", tc.CreateData)
}
