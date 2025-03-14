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

	for _, task := range tasks {
		if task.Id == id {
			return &task
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

	s, err := SerializeToJSON(tasks)

	if err != nil {
		fmt.Printf("unable to serialize task: %v", err)
	}

	SaveToJSON("db.json", s)

	return task
}

func UpdateTask(id int, description string) (models.Task) {
	tasks := GetTasks()

	task := GetTask(id)

	task.Description = description
	task.UpdatedAt = time.Now()

	s, err := SerializeToJSON(tasks)

	if err != nil {
		fmt.Printf("unable to serialize task: %v", err)
	}

	SaveToJSON("db.json", s)

	return *task
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
	
	s, err := SerializeToJSON(tasks)

	if err != nil {
		fmt.Printf("unable to serialize task: %v", err)
	}

	SaveToJSON("db.json", s)

	return *task
}