package main

import "fmt"

func main() {
	x := []int{17, 38, 57, 30, 27, 70}

	fmt.Println(bubbleSort(x))
}

func bubbleSort(num []int) []int {
	l := len(num)

	for i := 0; i < l; i++ {
		for j:=i+1; j < l; j++{
			if num[i] > num[i+1] {
				num[i], num[i+1] = num[j], num[i]
			}
		}
	}

	return num
}