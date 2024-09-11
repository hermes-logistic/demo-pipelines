package router

import (
	users_create_infrastructure "go-api/api/users/infrastructure/create"
	users_read_infrastructure "go-api/api/users/infrastructure/read"

	"github.com/gin-gonic/gin"
)

func (ro *Routes) UsersRoutes(r *gin.Engine, db interface{}) {
	ur := users_read_infrastructure.NewTaskHandler(db)
	r.GET("/users", ur.ReadData)

	uc := users_create_infrastructure.NewUserData(db)
	r.POST("/users", uc.CreateData)
}
