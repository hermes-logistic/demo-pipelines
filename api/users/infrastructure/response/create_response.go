package users_response_infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateConfirmResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Data created"})
}

// CreateErrorResponse responds with a generic error message.
func CreateErrorResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{"message": "Error"})
}

// CreateBadResponse responds with an error message indicating incorrect format.
func CreateBadResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "Format error"})
}
