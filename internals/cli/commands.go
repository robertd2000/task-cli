package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/robertd2000/task-cli/internals/models"
	"github.com/robertd2000/task-cli/internals/service"
)

type Commands struct{
	taskService service.TaskService
}

func NewCommands(taskService service.TaskService) *Commands {
	return &Commands{taskService: taskService}
}

func (c *Commands) Add(args []string) {
	if len(args) < 2 {
		fmt.Println("No description provided")
		return
	}

	description := args[1]

	c.taskService.CreateTask(description)

	fmt.Println("Task created")
}

func (c *Commands) Update(args []string) {
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
	
	c.taskService.UpdateTask(id,  &models.Task{Description: description})

	fmt.Printf("Task with id %d updated\n", id)
}

func (c *Commands) Delete(args []string) {
	if len(args) < 2 {
		log.Fatal("No id provided")
		return
	}
	idStr := args[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal("invalid id: %w", err)
	}

	c.taskService.DeleteTask(id)

	fmt.Printf("Task with id %d deleted\n", id)
}