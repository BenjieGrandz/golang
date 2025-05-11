package main

import "fmt"

func isValid(s string) bool {
	stack := []int{}

}

func main() {
    testCases := []string{
        "()",        // true
        "()[]{}",    // true
        "(]",        // false
        "([)]",      // false
        "{[]}",      // true
        "",          // true
        "{[}",       // false
        "([]{})",    // true
        "((((",      // false
        "(()())",    // true
    }

    for _, test := range testCases {
        fmt.Printf("Input: %q -> Valid: %v\n", test, isValid(test))
    }
}
