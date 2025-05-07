package main

import "fmt"

func arraySum(s []int) int {
	sum := 0

	for _, v := range s {
		sum += v
	}

	return sum
}

func main() {
	s := []int{1,2,3,4,5}
	fmt.Println(arraySum(s))
}