package tasks_delete_db_infrastructure

import (
	tasks_create_db_infrastructure "go-api/api/tasks/infrastructure/create/db"
	"go-api/logs"

	"gorm.io/gorm"
)

// DeleteTaskSql deletes a Task from the SQL database using gorm.
func DeleteTaskSql(db *gorm.DB, id string) bool {
	// Delete the Task with the specified ID from the corresponding table
	err := db.Delete(&tasks_create_db_infrastructure.SqlTask{}, id)

	if err != nil {
		// If there's an error during deletion, log the error in the error logger and exit
		logs.Error_Logger.Fatal(err)
		return false // Return false indicating deletion failed
	}

	// If deletion is successful, return true indicating deletion completed successfully
	return true
}
