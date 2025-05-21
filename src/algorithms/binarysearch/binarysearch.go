package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 7, 9, 11, 13}
	y := 9
	fmt.Printf("The target : %v is at index: %v", y, binaryS(x, y))
}

func binaryS(arr []int, target int) int {
	low := 0
	high := len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid -1
		}
	}

	return -1
}