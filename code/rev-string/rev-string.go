package main

import "fmt"

func revString(s string) string {
	rev := ""
	l := len(s)-1

	for i := l ; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}

// modify to include first word and last word functions
func main() {
	s := "Hello how are you?"
	fmt.Print(revString(s))
}
