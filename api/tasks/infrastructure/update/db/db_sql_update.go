package tasks_update_db_infrastructure

import (
	tasks_domain "go-api/api/tasks/domain"
	tasks_create_db_infrastructure "go-api/api/tasks/infrastructure/create/db"
	"go-api/logs"

	"gorm.io/gorm"
)

// UpdateTaskSql updates a task in the SQL database using gorm.
func UpdateTaskSql(db gorm.DB, data tasks_domain.Task) bool {
	// Try to update the record in the database
	err := db.Save(&tasks_create_db_infrastructure.SqlTask{
		ID:     data.ID,
		Name:   data.Name,
		Status: data.Status,
	}).Error

	if err != nil {
		// If there's an error updating, log the error in the error logger and return false
		logs.Error_Logger.Println("Error updating task in SQL:", err)
		return false // Return false indicating that the update failed
	}

	// If the update was successful, return true indicating that the update completed successfully
	return true
}
