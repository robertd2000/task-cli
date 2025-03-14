package main

import (
	"fmt"
	"os"

	"github.com/robertd2000/task-cli/internals/service"
)

func main() {
	args := os.Args[1:]

	fmt.Print(args)

	// service.CreateTask("test")

	// tasks:= service.GetTasks()

	// for _, task := range tasks {
	// 	fmt.Printf("%+v\n", task)
	// }

	// service.CreateTask("test 2")

	// for _, task := range tasks {
	// 	fmt.Printf("%+v\n", task)
	// }

	// service.CreateTask("test 45")

	service.UpdateTask(0, "test")

	// service.DeleteTask(4)
}