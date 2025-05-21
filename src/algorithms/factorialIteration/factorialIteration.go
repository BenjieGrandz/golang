package main

import "fmt"

func main() {
	fmt.Println(factIt(7))
}

func factIt(n int) int {
	for i := n; i >= 1; i-- {
		n = n*i
	}
	return n
}