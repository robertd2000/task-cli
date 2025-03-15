package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/robertd2000/task-cli/internals/repository"
	"github.com/robertd2000/task-cli/internals/service"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}

	taskRepository := repository.NewTaskRepository("db.json")
	taskService := service.NewTaskService(taskRepository)

	switch operation := args[0]; operation {
	case "add":
		if len(args) < 2 {
			fmt.Println("No description provided")
			return
		}

		description := args[1]

		taskService.CreateTask(description)
	case "update":
		if len(args) < 2 {
			fmt.Println("No id provided")
			return
		}

		if len(args) < 3 {
			fmt.Println("No description provided")
			return
		}

		idStr, description := args[1], args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Errorf("invalid id: %w", err)
		}
		
		taskService.UpdateTask(id, description)
	case "delete":
		if len(args) < 2 {
			fmt.Println("No id provided")
			return
		}
		idStr := args[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Errorf("invalid id: %w", err)
		}

		taskService.DeleteTask(id)
	case "list":
		fmt.Println("Get task")
	case "mark-in-progress":
		fmt.Println("Get task")
	case "mark-done":
		fmt.Println("Get task")
	default:
		fmt.Println("Invalid operation")
		return
	}
}