package users_response_infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateDeleteResponse responds with a confirmation message that the data has been deleted.
func CreateDeleteResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Data deleted"})
}

// CreateDeleteErrorResponse responds with a generic error message for the delete operation.
func CreateDeleteErrorResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{"message": "Error"})
}
