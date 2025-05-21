package main

import "fmt"

func main(){
	x := []int{3, 4, 7, 2, 8,15}
	fmt.Println(linearS(x, 2))
}

// implementing linear search
func linearS(n []int, target int) int {
	for _, v := range n {
		if v == target {
			return v
		}
	}

	return -1
}