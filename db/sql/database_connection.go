package sql_db

import (
	"go-api/logs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a connection with the database using the provided connection string.
func (d *Database) Connect(str string) {
	// Set the connection string
	dsn := str

	// Attempt to open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// In case of error, log the error in the error logger and return
		logs.Error_Logger.Println("Error connecting to the database:", err)
		return
	}

	// If there are no errors, the connection was successful
	logs.Info_Logger.Println("Successful connection")

	// Assign the connection object to the Database structure
	d.DB = db
}
