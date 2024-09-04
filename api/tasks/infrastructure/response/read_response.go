package tasks_response_infrastructure

import (
	tasks_domain "go-api/api/tasks/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReadResponse responds with a status code 200 (OK) and a JSON containing the task data.
func CreateReadResponse(c *gin.Context, d []tasks_domain.Task) {
	c.JSON(http.StatusOK, gin.H{"data": d})
}

// CreateReadErrorResponse responds with a status code 408 (Request Timeout) and a JSON containing a generic error message.
func CreateReadErrorResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{"message": "Error"})
}
