package main

import "fmt"

func largestNumber(s []int) int {
	for i := 0; i < len(s)-1; i++ {
		for j := i+1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]  // hot swap or bubble swap
			} 
		}
	}

	return s[len(s)-1]
}

func main() {
	myslice := []int{1, 3, 7, 4, 5}
	fmt.Println(largestNumber(myslice))
}