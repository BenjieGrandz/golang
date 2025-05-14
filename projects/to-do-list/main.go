package main

import (
	"fmt"
	"os"
	"strconv"
)

// Task reps one to do item with a desctription and done status
type Task struct {
	Description	string
	Done	bool
}

// slice that holds all our current tasks
// tasks := []Task{}
var tasks []Task

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

func addTask(args []string) {
	if len(args) < 1 {
		fmt.Println("Kindly add Description")
		return
	}

	// first argument is used as description
	desc := args[0]

	// Append new Task to slice
	// struct literal 
	// basically creates a new Task struct value assigning desc to the decription field leaving done field default(false)
	tasks = append(tasks, Task{Description: desc})

	fmt.Println("Added Task: ", desc)
}

func deleteTask(args []string){
	if len(args) < 1 {
		fmt.Println("Kindly add Description")
		return
	}

	// convert the args to an int (task number)
	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid Task Number")
		return
	}

	// store the desc for feedback before deletion
	desc := tasks[index-1].Description

	// remove task by slicing around it
	tasks = append(tasks[:index-1], tasks[index:]...)

	fmt.Println("Deleted task: ", desc)
}