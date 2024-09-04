package tasks_create_infrastructure

import (
	"bytes"
	"encoding/json"
	tasks_domain "go-api/api/tasks/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	var th TaskHandler

	r := gin.Default()

	r.POST("/tasks", th.CreateData)
	data := tasks_domain.Task{
		Name:   "test task",
		Status: "To Do",
	}

	value, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(value))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEmptyGetData(t *testing.T) {
	var th TaskHandler

	r := gin.Default()

	r.POST("/tasks", th.CreateData)
	data := tasks_domain.Task{
		Name:   "",
		Status: "",
	}

	value, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(value))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusRequestTimeout, w.Code)
}

func TestInvalidGetData(t *testing.T) {
	var th TaskHandler

	r := gin.Default()

	r.POST("/tasks", th.CreateData)
	data := []byte(`{"Name":"New Task", Status:to Do}`)

	value, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(value))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
