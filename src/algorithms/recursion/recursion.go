package main

import "fmt"

func main() {
	num1 := 7
	num2 := 9
	fmt.Printf("For %v the Factorial is: %v while the Fibonacci is: %v \n",num1 ,fact(num1), fib(num1))
	fmt.Printf("For %v the Factorial is: %v while the Fibonacci is: %v \n",num2, fact(num2), fib(num2))
}

func fact(n int) int {
	if n == 0 {
		return -1
	}

	return n * fact(n-1)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}