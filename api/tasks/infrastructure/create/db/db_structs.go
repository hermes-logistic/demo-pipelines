package tasks_create_db_infrastructure

import (
	"gorm.io/gorm"
)

type SqlTask struct {
	gorm.Model
	ID     string `gorm:"not null; primary_key"`
	Name   string `gorm:"not null" validate:"required"`
	Status string `gorm:"not null" validate:"required"`
}

type NoSqlTask struct {
	ID     string `bson:"ID,omitempty"`
	Name   string `bson:"name,omitempty"`
	Status string `bson:"status,omitempty"`
}
