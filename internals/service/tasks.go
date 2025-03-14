package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/robertd2000/task-cli/internals/models"
)

type autoInc struct {
    sync.Mutex
    id int
}

func (a *autoInc) ID() (id int) {
    a.Lock()
    defer a.Unlock()

    id = a.id
    a.id++
    return
}

var ai autoInc 

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

func CreateTask(description string) (models.Task) {
	task := models.Task{Id: ai.ID(), Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}



	return task
}