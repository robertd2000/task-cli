package main

import (
	"fmt"
	"os"

	"github.com/robertd2000/task-cli/internals/repository"
	"github.com/robertd2000/task-cli/internals/service"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}

	switch operation := args[0]; operation {
	case "add":
		fmt.Println("Create task")
	case "update":
		fmt.Println("Update task")
	case "delete":
		fmt.Println("Delete task")
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
	fmt.Print(args)
		taskRepository := repository.NewTaskRepository("db.json")
		taskService := service.NewTaskService(taskRepository)
	// taskService.CreateTask("test 7")
	taskService.UpdateTask(0, "test 0")
}