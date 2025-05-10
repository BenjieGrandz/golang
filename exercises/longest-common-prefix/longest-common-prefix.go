package main

import "fmt" 

func longestCommonPrefix (strs string) string {
	if len(strs) == 0 {
		return ""
	}

	// find the shortest prefix by looping through array to compare length of strings
	shortest := len(strs[0])
	for _, str := range strs {
		if shortest < len(str) {
			shortest = len(str)
		}
	}
}