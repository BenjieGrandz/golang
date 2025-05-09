package main

import "fmt"

func romanToInt(s string) int {
    symbols := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

	result := 0

	for i := 0; i < len(s) ; i++ {
		currentValue := symbols[rune(s[i])]

		if i+1 < len(s) {
			nextValue := symbols[rune(s[i+1])]

			if currentValue < nextValue {
				result -= currentValue
			} else {
				result += currentValue
			}
		} else {
			result += currentValue
		}
	}

	return result
}

func main() {
    // Test cases
    testCases := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
    for _, test := range testCases {
        fmt.Printf("%s = %d\n", test, romanToInt(test))
    }
}























