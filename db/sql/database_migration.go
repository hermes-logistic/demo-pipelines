package sql_db

import (
	tasks_create_db_infrastructure "go-api/api/tasks/infrastructure/create/db"
	"go-api/logs"
)

// Migration executes database migrations to create or update necessary tables.
func (d *Database) Migration() {
	// Get the database connection object from the Database structure
	db := d.DB

	// Execute migration to create or update the 'SqlTask' table
	err := db.AutoMigrate(&tasks_create_db_infrastructure.SqlTask{})

	if err != nil {
		// In case of error, log the error in the error logger and return
		logs.Error_Logger.Println("Error migrating tasks table", err)
		return
	}

	// If there are no errors, the migration was successful
	logs.Info_Logger.Println("Migration successful")
}
