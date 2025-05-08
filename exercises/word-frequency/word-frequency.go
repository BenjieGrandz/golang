package main

import (
	"fmt"
	"strings"
)

func wordFrequency(s string) map[string]int {
	// empty map i.e wordCOunts := {}
	wordCounts := make(map[string]int)

	// splits string to slice i.e words = ["go", "go"....]
	words := strings.Fields(s)

	// loop through each word and update map
	for _, word := range words {
		wordCounts[word]++
	}

	// output the map
	for word, count := range wordCounts{	// word is the key and count is the value
		fmt.Printf("%s: %d\n", word, count)
	}

	return wordCounts
}

func main() {
	s := "go go gophers are awesome go"
	fmt.Println(wordFrequency(s))
}

// map iteration is randomized thus the printf func