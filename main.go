package main

import (
	"github.com/robertd2000/task-cli/internals/cli"
	"github.com/robertd2000/task-cli/internals/repository"
	"github.com/robertd2000/task-cli/internals/service"
)

func main() {
	taskRepository := repository.NewTaskRepository("db.json")
	taskService := service.NewTaskService(taskRepository)

	cli.CLI(taskService)
}