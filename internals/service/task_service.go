package service

import (
	"github.com/robertd2000/task-cli/internals/models"
	"github.com/robertd2000/task-cli/internals/repository"
)

type TaskService interface {
	GetTask(id int) (*models.Task, error)
	GetTasks(filterStatus string) ([]models.Task, error)
	CreateTask(description string) (*models.Task, error)
	UpdateTask(id int, update *models.Task) (models.Task, error)
	DeleteTask(id int) (models.Task, error)
}

type taskService struct {
	repository repository.TaskRepository
}

func NewTaskService(repository repository.TaskRepository) TaskService {
	return &taskService{repository: repository}
}

func (s *taskService) GetTask(id int) (*models.Task, error) {
	return s.repository.GetTask(id)
}

func (s *taskService) GetTasks(filterStatus string) ([]models.Task, error) {
	return s.repository.GetTasks(filterStatus)
}

func (s *taskService) CreateTask(description string) (*models.Task, error){
	return s.repository.CreateTask(description)
}

func (s *taskService) UpdateTask(id int, update *models.Task) (models.Task, error) {
	return s.repository.UpdateTask(id, update)
}

func (s *taskService) DeleteTask(id int) (models.Task, error) {
	return s.repository.DeleteTask(id)
}