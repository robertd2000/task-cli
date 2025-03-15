package repository

import "github.com/robertd2000/task-cli/internals/models"

type TaskRepository interface {
	GetTasks() ([]models.Task, error)
	CreateTask(description string) (models.Task, error)
	UpdateTask(id int, description string) (models.Task, error)
	DeleteTask(id int) (models.Task, error)
}