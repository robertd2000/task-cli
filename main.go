package main

import (
	"fmt"
	"os"
	"time"

	"github.com/robertd2000/task-cli/internals/models"
	"github.com/robertd2000/task-cli/internals/service"
)

func main() {
	args := os.Args[1:]

	fmt.Print(args)

	task := &models.Task{
		Id:          1,
		Description: args[0],
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks := []models.Task{*task}
	
	s, err := service.SerializeToJSON(tasks)

	if err != nil {
		fmt.Printf("unable to serialize task: %w", err)
	}

	err = service.SaveToJSON("db.json", s)

	if err != nil {
		fmt.Printf("unable to save to json: %w", err)
	}

	stream, err := service.ReadFromJSON("db.json")

	if err != nil {
		fmt.Printf("unable to read from json: %w", err)
	}

	t, err := service.DeserializeFromJSON[[]models.Task](stream)

	if err != nil {
		fmt.Printf("unable to deserialize task: %w", err)
	}

	fmt.Println(t)

	t2 := &models.Task{
		Id:          2,
		Description: "ccc",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, *t2)

	s, err = service.SerializeToJSON(tasks)

	if err != nil {
		fmt.Printf("unable to serialize task: %w", err)
	}

	service.SaveToJSON("db.json", s)
}