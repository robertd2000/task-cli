package cli

import (
	"fmt"
	"os"

	"github.com/robertd2000/task-cli/internals/service"
)

func CLI(taskService service.TaskService) {
	commands := NewCommands(taskService)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}

	command := args[0]

	switch command {
	case "add":
		commands.Add(args)
	case "update":
		commands.Update(args)
	case "delete":
		commands.Delete(args)
	case "list":
		commands.List(args)
	case "mark-done":
		commands.ChangeStatus(args, "done")
	case "mark-in-progress":
		commands.ChangeStatus(args, "in-progress")
	case "mark-todo":
		commands.ChangeStatus(args, "todo")
	default:
		fmt.Println("Invalid operation")
		return
	}
}