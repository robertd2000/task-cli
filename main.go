package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/robertd2000/task-cli/internals/models"
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
			log.Fatal("invalid id: %w", err)
		}
		
		taskService.UpdateTask(id,  &models.Task{Description: description})
	case "delete":
		if len(args) < 2 {
			log.Fatal("No id provided")
			return
		}
		idStr := args[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal("invalid id: %w", err)
		}

		taskService.DeleteTask(id)

		fmt.Printf("Task with id %d deleted\n", id)
	case "list":
		if len(args) < 1 {
			fmt.Println("No id provided")
			return
		}

		var tasks []models.Task
		var err error

		if (len(args) == 2) {
			filterStatus := args[1]

			tasks, err = taskService.GetTasks(filterStatus)
			if err != nil {
				log.Fatal("unable to get tasks: %w", err)
			}
		} else {
			tasks, err = taskService.GetTasks("all")
			if err != nil {
				log.Fatal("unable to get tasks: %w", err)
			}
		}

		for _, task := range tasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("No id provided")
			return
		}

		idStr := args[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal("invalid id: %w", err)
		}
		
		taskService.UpdateTask(id,  &models.Task{Status: "in-progress"})
	case "mark-done":
		if len(args) < 2 {
			fmt.Println("No id provided")
			return
		}

		idStr := args[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal("invalid id: %w", err)
		}
		
		taskService.UpdateTask(id,  &models.Task{Status: "done"})
	default:
		fmt.Println("Invalid operation")
		return
	}
}