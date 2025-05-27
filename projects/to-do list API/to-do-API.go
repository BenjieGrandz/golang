package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
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
	// encode the todos slice into json and send as resoinse

	json.NewEncoder(w).Encode(todos)
}