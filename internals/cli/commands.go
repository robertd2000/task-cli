package cli

import (
	"fmt"

	"github.com/robertd2000/task-cli/internals/service"
)

type Commands struct{
	taskService service.TaskService
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