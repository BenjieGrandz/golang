package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Please enter : [add/sub/div/multi] [num1] [num2]")
		return
	}
	
	op := os.Args[1]
	num1 := os.Args[2]
	num2 := os.Args[3]

	a, err1 := strconv.ParseFloat(num1, 64)
	b, err2 := strconv.ParseFloat(num2, 64)

	if err1 != nil && err2 != nil {
		fmt.Println("Error : both arguments must be integers")
		return
	}

	switch op {
	case "add" :
		fmt.Println("Result: ", a+b)
	case "sub":
		fmt.Println("Result: ", a-b)
	case "div":
		if b <= 0 {
			fmt.Println("Cannot divide by zero or negatives") 
		}
		fmt.Println("Result: ", a/b)
	case "multi":
		fmt.Println("Result: ", a*b)
	default :
		fmt.Println("Unknown Operation [add/div/multi/sub]")
	}
	
}