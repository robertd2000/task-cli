package cli

import (
	"fmt"

	"github.com/robertd2000/task-cli/internals/models"
)

func DisplayTasks(tasks []models.Task) {
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
			task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}