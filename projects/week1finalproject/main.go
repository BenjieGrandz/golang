package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

// sorting algorithms
func bubbleSort(arr []int) []int {
	n := len(array)

	for i:=0; i<n-1; i++ {
		swapped := false
		
		for j:=0; j<n-i-1; j++ {
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

func main() {
	nums := []int{}

	// read user input for values to be sorted
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter numbers: ")
	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)

	// choose sorting algorithm
	selectAlgorithm := bufio.NewReader(os.Stdin)
	fmt.Println("Enter numbers: ")
	s, _ := selectAlgorithm.ReadString('\n')
	s = strings.TrimSpace(s)

	switch s {
	case 1:
		fmt.Println("This is a BubbleSort: ", bubbleSort(nums))
	case 2:
		fmt.Println("This is a BubbleSort: ", insertionSort(nums))
	case 3:
		fmt.Println("This is a BubbleSort: ", mergeSort(nums))
	default:
		fmt.Println("ERROR: You have not chosen a sorting algorithm!!")
	}

}