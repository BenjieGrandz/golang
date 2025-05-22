package main

import (
	"fmt"
	"strings"
	"bufio" // for full line input
	"os"
	//"sort"
)

type Books struct {
	ID int
	Title string
	Author string
}

func searchByTitle(book []Books, Title string) *Books {
	for _, b := range book {
		if strings.EqualFold(b.Title, Title){ // case sensitive
			return &b
		}
	}

	return nil
}

func main() {
	book := []Books{
		{ID: 3, Title: "Golang Practice", Author: "Ben Koimett"},
		{ID: 2, Title: "Clean code", Author: "Robert C. Martin "},
		{ID: 1, Title: "The Pragmatic Programmer", Author: "Andrew Hunt"},
	}

	// handle user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Title: ")
	t, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Please Enter Title!", err)
		return
	}
	t = strings.TrimSpace(t)

	resultByTitle := searchByTitle(book, t)
	if resultByTitle != nil {
		fmt.Println("Book found by Title: ", *resultByTitle)
	} else {
		fmt.Println("Book not found by Title.")
	}
}
