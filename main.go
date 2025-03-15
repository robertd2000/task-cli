package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	fmt.Print(args)

	// taskService.CreateTask("test 7")
}