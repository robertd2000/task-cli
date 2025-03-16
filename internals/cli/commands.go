package cli

import (
	"fmt"
	"log"

	"github.com/robertd2000/task-cli/internals/models"
	"github.com/robertd2000/task-cli/internals/service"
	"github.com/robertd2000/task-cli/internals/utils"
)

type Commands struct{
	taskService service.TaskService
}

func NewCommands(taskService service.TaskService) *Commands {
	return &Commands{taskService: taskService}
}

func (c *Commands) Add(args []string) {
	description := utils.GetDescription(args, 1)

	_, err := c.taskService.CreateTask(description)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task with description %s created\n", description)
}

func (c *Commands) Update(args []string) {
	id := utils.GetId(args, 1)
	description := utils.GetDescription(args, 2)
	
	c.taskService.UpdateTask(id,  &models.Task{Description: description})

	fmt.Printf("Task with id %d updated\n", id)
}

func (c *Commands) Delete(args []string) {
	id := utils.GetId(args, 1)

	c.taskService.DeleteTask(id)

	fmt.Printf("Task with id %d deleted\n", id)
}

func (c *Commands) List(args []string) {
	var filterStatus string

	switch len(args) {
	case 1:
		filterStatus = "all"
	case 2:
		filterStatus = args[1]
	default:
		fmt.Println("Invalid arguments")
		return
	}

	tasks, err := c.taskService.GetTasks(filterStatus)
	if err != nil {
		log.Fatalf("unable to get tasks: %v", err)
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
			task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

func (c *Commands) ChangeStatus(args []string, status string) {
	id := utils.GetId(args, 1)
	
	c.taskService.UpdateTask(id,  &models.Task{Status: status})

	fmt.Printf("Task with id %d updated\n", id)
}