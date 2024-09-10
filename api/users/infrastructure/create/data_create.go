package users_create_infrastructure

import (
	"fmt"
	users_empty_application "go-api/api/users/application/empty"
	users_domain "go-api/api/users/domain"
	users_create_db_infrastructure "go-api/api/users/infrastructure/create/db"
	users_response_infrastructure "go-api/api/users/infrastructure/response"
	"go-api/logs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserHandler struct {
	db interface{}
}

func NewUserData(db interface{}) users_domain.CreateUserInterface {
	return &UserHandler{db: db}
}

func (uh UserHandler) CreateData(c *gin.Context) {
	namespace := uuid.New()
	name := []byte("users")
	id := uuid.NewSHA1(namespace, name)

	var post users_domain.Users

	if err := c.BindJSON(&post); err != nil {
		users_response_infrastructure.CreateBadResponse(c)
		logs.Warning_Logger.Println("Format error")
		return
	}

	validate := users_empty_application.EmptyValidation(&post)
	if !validate {
		users_response_infrastructure.CreateErrorResponse(c)
		logs.Warning_Logger.Println("Empty fields")
		return
	}

	typeToEvaluate := fmt.Sprintf("%T", uh.db)
	fmt.Println("Type to evaluate: ", typeToEvaluate)

	switch db := uh.db.(type) {
	case *gorm.DB:
		insertData := users_create_db_infrastructure.SqlUser{
			ID:       id.String(),
			Username: post.Username,
			Password: post.Password,
		}

		insert := users_create_db_infrastructure.InsertUser(*db, insertData)

		if !insert {
			users_response_infrastructure.CreateErrorResponse(c)
			return
		}

		users_response_infrastructure.CreateConfirmResponse(c)
	default:
		logs.Info_Logger.Println("Unsupported database")
	}
}
