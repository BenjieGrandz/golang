package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int{
	words := strings.Fields(s)

	for index, word := range words{
		if index == len(words)-1{
			return len(word)
		}
	}
	return -1
}

func main(){
	str := "  fly me	to the moon		"
	fmt.Println(lengthOfLastWord(str))
}