package tasks_read_db_infrastructure

import (
	tasks_domain "go-api/api/tasks/domain"
	tasks_create_db_infrastructure "go-api/api/tasks/infrastructure/create/db"
	"go-api/logs"

	"gorm.io/gorm"
)

// GetAllTaskSql gets all tasks from the SQL database using gorm.
func GetAllTaskSql(db *gorm.DB) ([]tasks_domain.Task, interface{}) {
	// Variable to store retrieved data from the database
	var data []tasks_create_db_infrastructure.SqlTask

	// Perform a query to get all tasks from the corresponding table
	if err := db.Find(&data).Error; err != nil {
		// If there's an error retrieving data, log the error in the error logger and return the error
		logs.Error_Logger.Println("Error fetching SQL tasks:", err)
		return nil, err
	}

	// Create a list of domain tasks from the data retrieved from the database
	var tasks []tasks_domain.Task
	for _, newData := range data {
		// For each retrieved record, create a new domain task
		task := tasks_domain.Task{
			ID:     newData.ID,
			Name:   newData.Name,
			Status: newData.Status,
		}
		// Add the created task to the tasks list
		tasks = append(tasks, task)
	}

	// Return the list of tasks and nil, indicating no error occurred
	return tasks, nil
}
