package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune {
		')' : '(',
		']' : '[',
		'}' : '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{' :
			stack = append(stack, char)  // push
		case ')', ']', '}' :
			if len(stack) <= 0 || stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	
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
