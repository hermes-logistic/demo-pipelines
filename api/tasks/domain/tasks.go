package tasks_domain

type Task struct {
	ID     string
	Name   string `validate:"required"`
	Status string `validate:"required"`
}
