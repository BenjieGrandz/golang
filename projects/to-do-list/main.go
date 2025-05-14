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
// var tasks []Task
tasks := []string{}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please add a command")
	}
}