package main

import (
	"fmt"
	//"sort"
)

type Books struct {
	ID int
	title string
	author string
}

func main() {
	book := []Books{
		{ID: 3, title: "Golang Practice", author: "Ben Koimett"},
		{ID: 2, title: "Clean code", author: "Robert C. Martin "},
		{ID: 1, title: "The Pragmatic Programmer", author: "Andrew Hunt"},
	}

	t:=""
	fmt.Println("Enter title: ")
	fmt.Scan(&t)

	fmt.Println(searchByTitle(book, t))
}

func searchByTitle(book []Books, title string) *Books {
	for _, b := range book {
		if b.title == title {
			return &b
		}
	}

	return nil
}