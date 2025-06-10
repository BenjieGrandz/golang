package main

import "fmt"

func searchInsertPosition(n []int, target int) int{
	i:=0

	for _, nums := range n{
		if nums >= target{
			return i
		}
		i++
	}
	return -1
}

func main(){
	arr := []int{1,3,5,8,13} 
	target := 2

	fmt.Println(searchInsertPosition(arr, target))
}