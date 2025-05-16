package main

import "fmt"

func main() {
	x := []int{17, 38, 57, 30, 27, 70}

	fmt.Println(bubbleSort(x))
}

func bubbleSort(num []int) []int {
	l := len(num)

	for i := range l {
		for j := i + 1; j < l; j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}

	return num
}
