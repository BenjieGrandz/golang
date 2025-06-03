package main

import (
	"fmt"
)

func main() {
	var myNumber int
	fmt.Println("Enter number to be converted: ")
	fmt.Scan(&myNumber)

	binary := ""
	for i:=myNumber; i>=1; i=i/2 {
		if i % 2 == 0 {
			binary = "0" + binary
		} else {
			binary = "1" + binary
		}
	}

	fmt.Printf("Then binary of %v is:\n", myNumber)
	fmt.Println(binary)
}