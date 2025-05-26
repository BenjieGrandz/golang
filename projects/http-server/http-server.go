package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request ) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "pong")
	} else {
		http.Error(w, "Method Not Allowed ", http.StatusMethodNotAllowed )
	}
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	fmt.Println("Server is running on localhost:8080 ...")
	http.ListenAndServe(":8080", nil)
}