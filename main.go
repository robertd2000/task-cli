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
	
	s, err := service.SerializeToJSON(task)

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

	t, err := service.DeserializeFromJSON[models.Task](stream)

	if err != nil {
		fmt.Printf("unable to deserialize task: %w", err)
	}

	fmt.Println(t)

	t.Status = "done2"

	s, err = service.SerializeToJSON(t)

	if err != nil {
		fmt.Printf("unable to serialize task: %w", err)
	}

	service.SaveToJSON("db.json", s)
}