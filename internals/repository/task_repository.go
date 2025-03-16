package repository

import (
	"fmt"
	"time"

	"slices"

	"github.com/robertd2000/task-cli/internals/models"
	"github.com/robertd2000/task-cli/internals/utils"
)

type TaskRepository interface {
	GetTasks(filterStatus string) ([]models.Task, error)
	GetTask(id int) (*models.Task, error)
	CreateTask(description string) (*models.Task, error)
	UpdateTask(id int, update *models.Task) (models.Task, error)
	DeleteTask(id int) (models.Task, error)
}

type taskRepository struct {
	sourceFile string
}
func NewTaskRepository(sourceFile string) TaskRepository {
	return &taskRepository{sourceFile: sourceFile}
}

func (r *taskRepository) GetTasks(filterStatus string) ([]models.Task, error) {
	stream, err := utils.ReadFromJSON(r.sourceFile)

	if err != nil {
		return []models.Task{}, nil
	}

	tasks, err := utils.DeserializeFromJSON[[]models.Task](stream)

	if err != nil {
		return nil, fmt.Errorf("unable to deserialize tasks: %w", err)
	}

	if filterStatus == "all" {
		return tasks, nil
	}

	tasks = utils.Filter(tasks, func(task models.Task) bool {
		return task.Status ==filterStatus
	})
	
	return tasks, nil
}

func (r *taskRepository) GetTask(id int) (*models.Task, error) {
	tasks, err := r.GetTasks("all")
	if err != nil {
		return nil, fmt.Errorf("unable to get tasks: %w", err)
	}

	for i := range tasks {
		if tasks[i].Id == id {
			return &tasks[i], nil
		}
	}

	return nil, fmt.Errorf("task with id %d not found", id)
}

func (r *taskRepository) newTask(description string) models.Task {
	tasks, err := r.GetTasks("all")
	if err != nil {
		return models.Task{}
	}

	var nextID int

	if len(tasks) == 0 {
		nextID = 1
	} else {
		previousTask := tasks[len(tasks)-1]
		nextID = previousTask.Id + 1
	}

	return *models.NewTask(nextID, description, "todo", time.Now(), time.Now())
}

func (r *taskRepository) CreateTask(description string) (*models.Task, error) {
	tasks, err := r.GetTasks("all")

	if err != nil {
		return nil, fmt.Errorf("unable to get tasks: %w", err)
	}

	task := r.newTask(description)
	tasks = append(tasks, task)

	r.commit(tasks)

	return &task, nil
}

func (r *taskRepository) UpdateTask(id int, update *models.Task) (models.Task, error) {
	tasks, err := r.GetTasks("all")
	if err != nil {
		return models.Task{}, fmt.Errorf("unable to get tasks: %w", err)
	}

	if tasks == nil {
		return models.Task{}, fmt.Errorf("no tasks found")
	}

	var task *models.Task
	for i := range tasks {
		if tasks[i].Id == id {
			if update.Description != "" {
                tasks[i].Description = update.Description
            }
            if update.Status != "" {
                tasks[i].Status = update.Status
            }
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

func (r *taskRepository) DeleteTask(taskId int) (models.Task, error) {
	tasks, err := r.GetTasks("all")
	if err != nil {
		return models.Task{}, fmt.Errorf("unable to get tasks: %w", err)
	}

	task, err := r.GetTask(taskId)
	if err != nil {
		return models.Task{}, fmt.Errorf("task with id %d not found", taskId)
	}
	
	id := task.Id

	for i, task := range tasks {
		if task.Id == id {
			tasks = slices.Delete(tasks, i, i+1)
		}
	}
	
	r.commit(tasks)

	return *task, nil
}

func (r *taskRepository) commit(tasks []models.Task) error {
	s, err := utils.SerializeToJSON(tasks)
	if err != nil {
		return fmt.Errorf("unable to serialize task: %w", err)
	}

	if err := utils.SaveToJSON(r.sourceFile, s); err != nil {
		return fmt.Errorf("unable to save to JSON: %w", err)
	}

	return nil
}
