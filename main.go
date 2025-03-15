package main

import (
	"fmt"
	"os"

	"github.com/robertd2000/task-cli/internals/repository"
	"github.com/robertd2000/task-cli/internals/service"
)

func main() {
	args := os.Args[1:]

	fmt.Print(args)
		taskRepository := repository.NewTaskRepository("db.json")
		taskService := service.NewTaskService(taskRepository)
	// taskService.CreateTask("test 7")
	taskService.UpdateTask(0, "test 0")
}