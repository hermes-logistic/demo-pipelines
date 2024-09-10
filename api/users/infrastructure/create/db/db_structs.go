package users_create_db_infrastructure

import "gorm.io/gorm"

type SqlUser struct {
	gorm.Model
	ID       string `gorm:"not null; primary key"`
	Username string `gorm:"not null" validate:"required"`
	Password string `gorm:"not null" validate:"required"`
}
