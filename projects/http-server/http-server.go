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

}