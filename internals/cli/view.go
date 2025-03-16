package cli

import (
	"fmt"

	"github.com/robertd2000/task-cli/internals/models"
)

func DisplayTasks(tasks []models.Task) {
	for _, task := range tasks {
		fmt.Printf("%s\n", task.Display())
	}
}