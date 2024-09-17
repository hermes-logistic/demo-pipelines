package users_response_infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUpdateResponse responds with a confirmation message that the data has been updated.
func CreateUpdateResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Data updated"})
}

// CreateUpdateErrorResponse responds with a generic error message for the update operation.
func CreateUpdateErrorResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{"message": "Error"})
}
