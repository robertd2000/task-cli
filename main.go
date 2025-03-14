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
	
	s, err := service.SerializeTask(task)

	if err != nil {
		fmt.Printf("unable to serialize task: %w", err)
	}

	err = service.SaveToJSON(s)

	if err != nil {
		fmt.Printf("unable to save to json: %w", err)
	}
}