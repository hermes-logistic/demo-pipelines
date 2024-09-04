package tasks_create_db_infrastructure

import (
	"go-api/logs"

	"gorm.io/gorm"
)

// InsertTask inserts a new record into the "tasks" table using gorm.
func InsertTask(db gorm.DB, data SqlTask) bool {
	// Create the record in the "tasks" table using the db object
	insert := db.Create(&data)

	if insert.Error != nil {
		// If there's an error during insertion, log the error in the error logger
		logs.Error_Logger.Println("Error inserting data:", insert.Error)
		return false // Return false indicating insertion failed
	}

	// If insertion is successful, log an information message in the logger
	logs.Info_Logger.Println("Data inserted successfully")
	return true // Return true indicating insertion was successful
}
