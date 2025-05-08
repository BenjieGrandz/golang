package main

import "fmt"

func charFrequency(s string) map[string]int {
	wordCount := make(map[string]int)

	for _, char := range s {
		wordCount[string(char)]++
	}

	for char, count := range wordCount {
		fmt.Printf("%s: %d\n", char, count)
	}

	return wordCount
}

func main() {
	s := "hello"
	charFrequency(s)
}