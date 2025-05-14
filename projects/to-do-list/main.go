package main

import (
	"fmt"
	"os"
	"strconv"
)

// Task reps on to do item with a desctription and done status
type Task struct {
	Description	string
	Done	bool
}

// slice that holds all our current tasks
var tasks []Task
// tasks := []Task{}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please add a command and/or arguments")
		return
	}

	// first argument is command
	command := os.Args[1]

	// switch through commands
	switch command {
	case "add":
		addTask(os.Args[2:]) // check if just normal indexing works as well
	case "list":
		listTasks()
	case "done":
		markDone(os.Args[2:])
	case "delete":
		fmt.Println("Unknown Command", command)
	}
}

// add functions her ordered logically and by importance

func addTask(args string) {
	if len(args) < 1 {
		fmt.Println("Kindly add Description")
	}

	// first argument is used as description
	desc := string(args[0])

	// Append new Task to slice
	// struct literal 
	// basically creates a new Task struct value assigning desc to the decription field leaving done field default(false)
	tasks = append(tasks, Task{Description: desc})

	fmt.Println("Added Task: ", desc)
}