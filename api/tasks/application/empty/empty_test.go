package tasks_empty_application

import (
	tasks_domain "go-api/api/tasks/domain"
	"testing"
)

func TestEmptyValidation(t *testing.T) {
	task := tasks_domain.Task{
		Name:   "Task 1",
		Status: "Pending",
	}
	if result := EmptyValidation(&task); !result {
		t.Error("Se esperaba que la validación pasara, pero falló")
	}

	emptyTask := tasks_domain.Task{}
	if result := EmptyValidation(&emptyTask); result {
		t.Error("Se esperaba que la validación fallara, pero pasó")
	}
}
