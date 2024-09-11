package users_response_infrastructure

import (
	users_domain "go-api/api/users/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateReadResponse(c *gin.Context, d []users_domain.Users) {
	c.JSON(http.StatusOK, gin.H{"data": d})
}

func CreateReadErrorResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{"message": "Error"})
}
