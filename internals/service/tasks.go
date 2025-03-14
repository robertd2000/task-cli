package service

import (
	"fmt"
	"time"

	"github.com/robertd2000/task-cli/internals/models"
)

func GetTasks() []models.Task {
	stream, err := ReadFromJSON("db.json")

	if err != nil {
		fmt.Printf("unable to read from json: %w", err)
	}

	tasks, err := DeserializeFromJSON[[]models.Task](stream)

	if err != nil {
		fmt.Printf("unable to deserialize task: %w", err)
	}

	return tasks
}

func GetTask(id int) (models.Task) {
	tasks := GetTasks()

	return tasks[id]
}

func CreateTask(description string) (models.Task) {

	tasks := GetTasks()
	prevTask := tasks[len(tasks)-1]
	id := prevTask.Id + 1
	task := models.Task{Id: id, Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	tasks = append(tasks, task)

	s, err := SerializeToJSON(tasks)

	if err != nil {
		fmt.Printf("unable to serialize task: %w", err)
	}

	SaveToJSON("db.json", s)

	return task
}

func UpdateTask(id int, description string) (models.Task) {
	tasks := GetTasks()

	task := &tasks[id]

	task.Description = description
	task.UpdatedAt = time.Now()

	s, err := SerializeToJSON(tasks)

	if err != nil {
		fmt.Printf("unable to serialize task: %w", err)
	}

	SaveToJSON("db.json", s)

	return *task
}

func DeleteTask(id int) (models.Task) {
	tasks := GetTasks()

	task := &tasks[id]

	tasks = append(tasks[:id], tasks[id+1:]...)

	s, err := SerializeToJSON(tasks)

	if err != nil {
		fmt.Printf("unable to serialize task: %w", err)
	}

	SaveToJSON("db.json", s)

	return *task
}