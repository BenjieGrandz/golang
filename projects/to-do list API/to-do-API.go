package main

import (
	"encoding/json"
	// "fmt"
	"net/http"
	// "strings"
	// "strconv"
	"sync"
)

type Todo struct {
	ID int `json: "ID"`
	Task string `json: "task"`
}

// declare shared variables

var (
	todos = []Todo{}
	nextID = 1
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

}

func main() {
	
}