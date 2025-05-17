package main

import "fmt"

func main() {
	x := []int{7, 6, 3, 9, 1, 2}
	fmt.Println(insertSort(x))
}

func insertSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		key := n[i]
		j := i - 1

		for j >= 0 && n[j] > key {
			n[j+1] = n[j]
			j--
		}

		n[j+1] = key
	}

	return n
}
