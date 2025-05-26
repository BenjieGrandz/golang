package main

import (
	"fmt"
	"net/http"
)

// Ping Handler FUnction
func pingHandler(w http.ResponseWriter, r *http.Request ) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "pong")
	} else {
		http.Error(w, "Method Not Allowed ", http.StatusMethodNotAllowed )
	}
}

// hello Handler Function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "hello")
	} else {
		http.Error(w, "Mehtod Not Allowed ", http.StatusMethodNotAllowed)
	}
}

// goodbye handler function
func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "goodbye")
	} else {
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/goodbye", goodbyeHandler)
	fmt.Println("Server is running on localhost:8080 ...")
	http.ListenAndServe(":8080", nil)
}