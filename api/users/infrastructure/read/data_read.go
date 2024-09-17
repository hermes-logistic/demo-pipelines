package users_read_infrastructure

import (
	users_domain "go-api/api/users/domain"
	users_read_db_infrastructure "go-api/api/users/infrastructure/read/db"
	users_response_infrastructure "go-api/api/users/infrastructure/response"
	"go-api/logs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db interface{}
}

func NewTaskHandler(db interface{}) users_domain.ReadUserInterface {
	return &UserHandler{db: db}
}

func (uh *UserHandler) ReadData(c *gin.Context) {
	switch db := uh.db.(type) {
	case *gorm.DB:
		users, err := users_read_db_infrastructure.GetAllUsersSql(db)

		if err != nil {
			users_response_infrastructure.CreateReadErrorResponse(c)
			return
		}

		users_response_infrastructure.CreateReadResponse(c, users)
		return
	default:
		logs.Info_Logger.Println("No support for this database")
	}

}
