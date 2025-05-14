package main

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	tasks = []Task{} // reset tasks

	addTask([]string{"Test Task"})
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Description != "Test Task" {
		t.Errorf("Expected 'Test Task', got '%s'", tasks[0].Description)
	}
	if tasks[0].Done {
		t.Errorf("Expected task to be not done")
	}
}

func TestMarkDone(t *testing.T) {
	tasks = []Task{
		{Description: "Incomplete Task", Done: false},
	}

	markDone([]string{"1"})
	if !tasks[0].Done {
		t.Errorf("Expected task to be marked as done")
	}
}

func TestDeleteTask(t *testing.T) {
	tasks = []Task{
		{Description: "Task 1", Done: false},
		{Description: "Task 2", Done: false},
	}

	deleteTask([]string{"1"})
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task after delete, got %d", len(tasks))
	}
	if tasks[0].Description != "Task 2" {
		t.Errorf("Expected remaining task to be 'Task 2', got '%s'", tasks[0].Description)
	}
}

func TestListTasksEmpty(t *testing.T) {
	tasks = []Task{} // reset
	listTasks()      // should print "No Tasks Available"
}
