package main

import "fmt"

func mergeSort(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr)/2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main(){
	num := []ibt{2,8,3,7,4,6,5}
	fmt.Println(mergeSort(num))
}