package main

import "fmt"

func main() {
	x := []int{17, 38, 57, 30, 27, 70}

	fmt.Println(bubbleSort(x))
}

func bubbleSort(num []int) []int {
	n := len(num)

	for i:=0; i<n-1; i++ {
		for j:=0; j<n-i-1;j++ {
			swapped := true

			if num[j] > num[J+1] {
				num[j], num[j+1] = num[j+1],num[j]
				swapped = false
			}

			if !swapped {
				break
			}
		}
	}
	return num
}
