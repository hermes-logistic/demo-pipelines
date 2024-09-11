package users_read_db_infrastructure

import (
	users_domain "go-api/api/users/domain"
	users_create_db_infrastructure "go-api/api/users/infrastructure/create/db"
	"go-api/logs"

	"gorm.io/gorm"
)

func GetAllUsersSql(db *gorm.DB) ([]users_domain.Users, interface{}) {
	var data []users_create_db_infrastructure.SqlUser

	if err := db.Find(&data).Error; err != nil {
		logs.Error_Logger.Println("Error fetching SQL Users:", err)
		return nil, err
	}

	var users []users_domain.Users

	for _, newData := range data {
		user := users_domain.Users{
			ID:       newData.ID,
			Username: newData.Username,
			Password: newData.Password,
		}
		users = append(users, user)
	}

	return users, nil
}
