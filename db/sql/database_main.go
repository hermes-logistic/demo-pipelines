package sql_db

import "gorm.io/gorm"

// Data structure that holds the context of the database
type Database struct {
	DB *gorm.DB
}
