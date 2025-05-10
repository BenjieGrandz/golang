package main

import "fmt" 

func longestCommonPrefix (strs string) string {
	if len(strs) == 0 {
		return ""
	}

	// find the shortest prefix by looping through array to compare length of strings
	shortest := len(strs[0])
	for _, str := range strs {
		if shortest > len(str) {
			shortest = len(str)
		}
	}

	for i := 0; i < shortest ; i++ {
		target := strs[0][i]

		for _, char := range strs {
			if str[i] != target {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:shortest]
}

func main() {
    // Test cases from examples
    testCases := [][]string{
        {"flower", "flow", "flight"},
        {"dog", "racecar", "car"},
        {},
        {"a"},
        {"ab", "abc", "abcd"},
    }
    
    for _, testCase := range testCases {
        fmt.Printf("Input: %v\nOutput: %q\n\n", testCase, longestCommonPrefix(testCase))
    }
}