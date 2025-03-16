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

	_, err := c.taskService.CreateTask(description)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task with description %s created\n", description)
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

func (c *Commands) List(args []string) {
	if len(args) < 1 {
		fmt.Println("No id provided")
		return
	}

	var tasks []models.Task
	var err error

	if (len(args) == 2) {
		filterStatus := args[1]

		tasks, err = c.taskService.GetTasks(filterStatus)
		if err != nil {
			log.Fatal("unable to get tasks: %w", err)
		}
	} else {
		tasks, err = c.taskService.GetTasks("all")
		if err != nil {
			log.Fatal("unable to get tasks: %w", err)
		}
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

func (c *Commands) ChangeStatus(args []string, status string) {
	if len(args) < 2 {
		fmt.Println("No id provided")
		return
	}

	idStr := args[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal("invalid id: %w", err)
	}
	
	c.taskService.UpdateTask(id,  &models.Task{Status: status})

	fmt.Printf("Task with id %d updated\n", id)
}