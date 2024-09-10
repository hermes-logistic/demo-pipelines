package router

import (
	users_create_infrastructure "go-api/api/users/infrastructure/create"

	"github.com/gin-gonic/gin"
)

func (ro *Routes) UsersRoutes(r *gin.Engine, db interface{}) {
	ur := users_create_infrastructure.NewUserData(db)
	r.POST("/users", ur.CreateData)
}
