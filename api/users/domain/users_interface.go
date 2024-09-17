package users_domain

import "github.com/gin-gonic/gin"

type UserInterface interface {
}

type CreateUserInterface interface {
	CreateData(c *gin.Context)
}

type ReadUserInterface interface {
	ReadData(c *gin.Context)
}

type UpdateUserInterface interface {
	UpdateData(c *gin.Context)
}

type DeleteUserInterface interface {
	DeleteData(c *gin.Context)
}
