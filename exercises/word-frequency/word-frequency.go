package main

import (
	"fmt"
	"strings"
)

func wordFrequency(s string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		fmt.Println(word, count)
	}

	return wordCount
}

func main() {
	s := "go go gophers are awesome go"
	wordFrequency(s)
}

// map iteration is randomized thus the printf func