package main

import "fmt"

type books struct {
	ID int
	title string
	author string
}

func main() {
	books := []books{
		{ID: 3, title: "Golang Practice", author: "Ben Koimett"},
		{ID: 2, title: "Clean code", author: "Robert C. Martin "},
		{ID: 1, title: "The Pragmatic Programmer", author: "Andrew Hunt"},
	}
}