package service

import (
	"github.com/robertd2000/task-cli/internals/models"
	"github.com/robertd2000/task-cli/internals/repository"
)

type TaskService struct {
	repository repository.ITaskRepository
}

func (s *TaskService) GetTask(id int) (*models.Task, error) {
	return s.repository.GetTask(id)
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.repository.GetTasks()
}

func (s *TaskService) CreateTask(description string) (*models.Task, error){
	return s.repository.CreateTask(description)
}

func (s *TaskService) UpdateTask(id int, description string) (models.Task, error) {
	return s.repository.UpdateTask(id, description)
}

func (s *TaskService) DeleteTask(id int) (models.Task, error) {
	return s.DeleteTask(id)
}