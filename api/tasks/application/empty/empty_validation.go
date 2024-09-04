package tasks_empty_application

import (
	tasks_domain "go-api/api/tasks/domain"
	structs_helpers "go-api/helpers/structs"
)

// EmptyValidation performs validation to determine if a Task has empty or null fields.
func EmptyValidation(t *tasks_domain.Task) bool {
	// Create a new instance of Task with the same values as the original Task
	newTask := tasks_domain.Task{
		Name:   t.Name,
		Status: t.Status,
	}

	// Get the validator for the Task structure
	validate := structs_helpers.StructValidator

	// Validate the newTask structure using the validator
	response := structs_helpers.ValidateStruct(validate, newTask)

	// Check the validation response
	if !response {
		// If the validation fails (response is false), it means there are empty or null fields
		// Return false to indicate that the validation did not pass
		return false
	} else {
		// If the validation passes (response is true), it means there are no empty or null fields
		// Return true to indicate that the validation passed
		return true
	}
}
