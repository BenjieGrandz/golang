package main

import "fmt"

func plusOne(arr []int) []int{
	n := len(arr)

	for i:=n-1;i>=0;i--{
		if arr[i] < 9{
			arr[i]++
			return arr
		}
		arr[i]=0
	}
	return append([]int{1}, arr...)
}

func main(){
	x:=[]int{1,2,3}
	fmt.Println(plusOne(x))
}