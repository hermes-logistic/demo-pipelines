package users_create_db_infrastructure

import (
	"go-api/logs"

	"gorm.io/gorm"
)

func InsertUser(db gorm.DB, data SqlUser) bool {
	insert := db.Create(&data)

	if insert.Error != nil {
		logs.Error_Logger.Println("Error inserting data:", insert.Error)
		return false // Return false indicating insertion failed
	}

	logs.Info_Logger.Println("Data inserted successfully")
	return true // Return true indicating insertion was successful
}
