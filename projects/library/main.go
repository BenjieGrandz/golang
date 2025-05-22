package main

import (
	"bufio" // for full line input
	"fmt"
	"os"
	"sort"
	"strings"
)

type Books struct {
	ID     int
	Title  string
	Author string
}

func searchByTitle(book []Books, Title string) *Books {
	for _, b := range book {
		if strings.EqualFold(b.Title, Title) { // case sensitive
			return &b
		}
	}

	return nil
}

func binarySearchByID(book []Books, ID int) *Books {
	low, high := 0, len(book)-1

	for low <= high {
		mid := low + (high-low)/2
		if book[mid].ID == ID {
			return &book[mid]
		} else if book[mid].ID < ID {
			low = mid + 1
		} else {
			high = mid - 1
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

	// sort books first
	sort.Slice(book, func(i, j int) bool {
		return book[i].ID < book[j].ID
	})

	// compare the books by ID
	var IDinput = 0
	fmt.Println("Please enter ID of the book: ")
	fmt.Scan(&IDinput)

	resultByBinary := binarySearchByID(book, IDinput)
	if resultByBinary != nil {
		fmt.Println("Book found by ID: ", *resultByBinary)
	} else {
		fmt.Println("Book not found by ID")
	}
}
