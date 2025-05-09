package main

import "fmt"

func palindromeNumber(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	rev := 0
	for x > rev {
		digit := x%10
		rev = (rev*10) + digit
		x /= 10
	}

	return x == rev || x == rev/10
}

func main() {
    tests := []int{121, -121, 10, 12321, 0, 1221, 123}
    for _, num := range tests {
        fmt.Println(palindromeNumber(num))
    }
}
