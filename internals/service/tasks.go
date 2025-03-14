package service

import (
	"fmt"
	"time"

	"github.com/robertd2000/task-cli/internals/models"
)

func GetTasks() []models.Task {
	stream, err := ReadFromJSON("db.json")

	if err != nil {
		fmt.Printf("unable to read from json: %v", err)
	}

	tasks, err := DeserializeFromJSON[[]models.Task](stream)

	if err != nil {
		fmt.Printf("unable to deserialize task: %v", err)
	}

	return tasks
}

func GetTask(id int) (*models.Task) {
	tasks := GetTasks()

	for i := range tasks {
		if tasks[i].Id == id {
			return &tasks[i]
		}
	}

	return nil
}

func newTask(description string) (models.Task) {
	tasks := GetTasks()
	prevTask := tasks[len(tasks)-1]
	id := prevTask.Id + 1
	task := models.Task{Id: id, Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	return task
}

func CreateTask(description string) (models.Task) {
	tasks := GetTasks()
	task := newTask(description)
	tasks = append(tasks, task)

	commit(tasks)

	return task
}

func UpdateTask(id int, description string) (models.Task, error) {
	tasks := GetTasks()

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

	if err := commit(tasks); err != nil {
		return models.Task{}, fmt.Errorf("failed to commit changes: %v", err)
	}

	return *task, nil
}

func DeleteTask(taskId int) (models.Task) {
	tasks := GetTasks()

	task := GetTask(taskId)
	id := task.Id

	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	
	commit(tasks)

	return *task
}

func commit(tasks []models.Task) error {
	s, err := SerializeToJSON(tasks)
	if err != nil {
		return fmt.Errorf("unable to serialize task: %w", err)
	}

	if err := SaveToJSON("db.json", s); err != nil {
		return fmt.Errorf("unable to save to JSON: %w", err)
	}

	return nil
}
