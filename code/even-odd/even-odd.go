package main

import "fmt"

func evenOdd(n int) string {
	if n < 0 {
		return "This is a NEGATIVE number"
	}

	if n%2 == 0 {
		return "This is an EVEN number"
	}

	return "This is an ODD number"
}

func main() {
	fmt.Println(evenOdd(5)) // odd
	fmt.Println(evenOdd(4)) // even
	fmt.Println(evenOdd(-5)) // negative
}