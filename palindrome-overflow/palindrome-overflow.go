package main

import (
    "fmt"
)

func reverse(x int) int {
    if x < 0 || ( x%10 == 0 && x != 0) {
		return x
	}

	rev := 0
	for x > rev {
		rev = (rev*10) + x%10
		x /= 10
	}

}

func main() {
    tests := []int{123, -456, 120, 0, 1534236469, -2147483412}
    for _, num := range tests {
        fmt.Printf("reverse(%d) = %d\n", num, reverse(num))
    }
}




// [-2³¹, 2³¹ - 1]