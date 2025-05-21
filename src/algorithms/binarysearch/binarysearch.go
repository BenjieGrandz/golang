package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 7, 9, 11, 13}
	y := 9
	fmt.Printf("The target : %v is at index: %v", y, binaryS(x))
}