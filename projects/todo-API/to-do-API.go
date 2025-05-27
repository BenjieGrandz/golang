package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	//"log"
	"sync"
)

type Todo struct {
	ID   int    `json: "ID"`
	Task string `json: "task"`
}

// declare shared variables

var (
	todos   = []Todo{}
	nextID  = 1
	todoMux sync.Mutex
)

// Handler for GET/ todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	// encode the todos slice into json and send as response

	json.NewEncoder(w).Encode(todos)
}

// handler for POST/ todos
func createTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	// decode the JSOn body into a Todo struct
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // error 400
		return
	}

	//lock before modifying shared data
	todoMux.Lock()

	t.ID = nextID
	nextID++
	todos = append(todos, t)

	todoMux.Unlock()
	// unclock after modifulyin data

	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(t)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// get dodo ID
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// lock before modifying shared data
	todoMux.Lock()
	defer todoMux.Unlock()

	// search for todo by ID then remove by slicing slice
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // 204 - No content
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound) // if not found
}

func main() {
	// Route: handle GET and POST on /todos
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodos(w, r)
		case http.MethodPost:
			createTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route : handle DELETE on /todos/{id}
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			deleteTodo(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server on port 8080
	fmt.Println("To-DO server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
