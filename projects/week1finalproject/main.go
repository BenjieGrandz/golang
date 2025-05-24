package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

// sorting algorithms
func bubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}

// insertion sort
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j > 0 && arr[j] > key {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}

	return arr
}

// merge sort algorithm - divide and rule
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// merge - used to sort and combine
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {

	// read user input for values to be sorted
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter numbers: ")
	// n, _ := reader.ReadString('\n')
	// x := strconv.Atoi(n)

	// length of array
	var arrLen int
	fmt.Println("Enter the number of values to be sorted")
	fmt.Scan(&arrLen)

	valueToSort := 0
	nums := []int{}
	// new code to take in data for arrays to be sorted
	for i:=1; i<=arrLen;i++{
		fmt.Printf("Enter number %v:\n",i)
		fmt.Scan(&valueToSort)
		nums = append(nums, valueToSort)
	}

	// choose sorting algorithm
	selectAlgorithm := bufio.NewReader(os.Stdin)
	fmt.Println("Enter sorting type you want: ")
	s, _ := selectAlgorithm.ReadString('\n')
	s = strings.TrimSpace(s)

	switch s {
	case "A":
		fmt.Println("This is a BubbleSort: ", bubbleSort(nums))
	case "B":
		fmt.Println("This is a BubbleSort: ", insertionSort(nums))
	case "C":
		fmt.Println("This is a BubbleSort: ", mergeSort(nums))
	default:
		fmt.Println("ERROR: You have not chosen a sorting algorithm!!")
	}

}
