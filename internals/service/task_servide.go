package service

import (
	"github.com/robertd2000/task-cli/internals/models"
	"github.com/robertd2000/task-cli/internals/repository"
)

type TaskService struct {
	repository repository.ITaskRepository
}
func (s *TaskService) GetTask(id int) *models.Task {
	return s.GetTask(id)
}