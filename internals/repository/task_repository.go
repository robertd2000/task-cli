package repository

import (
	"fmt"
	"time"

	"github.com/robertd2000/task-cli/internals/models"
	"github.com/robertd2000/task-cli/internals/utils"
)

type ITaskRepository interface {
	GetTasks() ([]models.Task, error)
	GetTask(id int) (*models.Task, error)
	CreateTask(description string) (*models.Task, error)
	UpdateTask(id int, description string) (models.Task, error)
	DeleteTask(id int) (models.Task, error)
}

type TaskRepository struct {
	sourceFile string
}

func (r *TaskRepository) GetTasks() []models.Task {
	stream, err := utils.ReadFromJSON(r.sourceFile)

	if err != nil {
		fmt.Printf("unable to read from json: %v", err)
	}

	tasks, err := utils.DeserializeFromJSON[[]models.Task](stream)

	if err != nil {
		fmt.Printf("unable to deserialize task: %v", err)
	}

	return tasks
}

func (r *TaskRepository) GetTask(id int) (*models.Task) {
	tasks := r.GetTasks()

	for i := range tasks {
		if tasks[i].Id == id {
			return &tasks[i]
		}
	}

	return nil
}

func (r *TaskRepository) newTask(description string) (models.Task) {
	tasks := r.GetTasks()
	prevTask := tasks[len(tasks)-1]
	id := prevTask.Id + 1
	task := models.Task{Id: id, Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	return task
}

func (r *TaskRepository) CreateTask(description string) (*models.Task) {
	tasks := r.GetTasks()
	task := r.newTask(description)
	tasks = append(tasks, task)

	r.commit(tasks)

	return &task
}

func (r *TaskRepository) UpdateTask(id int, description string) (models.Task, error) {
	tasks := r.GetTasks()

	if tasks == nil {
		return models.Task{}, fmt.Errorf("no tasks found")
	}

	var task *models.Task
	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			task = &tasks[i]
			break
		}
	}

	if task == nil {
		return models.Task{}, fmt.Errorf("task with id %d not found", id)
	}

	if err := r.commit(tasks); err != nil {
		return models.Task{}, fmt.Errorf("failed to commit changes: %v", err)
	}

	return *task, nil
}

func (r *TaskRepository) DeleteTask(taskId int) (models.Task) {
	tasks := r.GetTasks()

	task := r.GetTask(taskId)
	id := task.Id

	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	
	r.commit(tasks)

	return *task
}

func (r *TaskRepository) commit(tasks []models.Task) error {
	s, err := utils.SerializeToJSON(tasks)
	if err != nil {
		return fmt.Errorf("unable to serialize task: %w", err)
	}

	if err := utils.SaveToJSON(r.sourceFile, s); err != nil {
		return fmt.Errorf("unable to save to JSON: %w", err)
	}

	return nil
}
