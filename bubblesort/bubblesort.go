package main

import "fmt"

func main() {
	x := []int{17, 38, 57, 30, 27, 70}

	for _, v := range x {
		fmt.Println(bubbleSort(x))
	}
}

func bubbleSort(num []int) []int {
	
}