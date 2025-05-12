package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	names := []string{"Alice", "Benjie", "Cytnthia", "Dennis", "Elvis"}

	// create the file
	file, err := os.Create("names.txt")
	if err != nil {
		fmt.Println("File creation error: ", err)
		return
	}
	defer file.Close()


}

// file, err := os.Create("file.txt")
// defer file.Close()

// file.WriteString("Hello\n")

// file, _ := os.Open("file.txt")
// scanner := bufio.NewScanner(file)
// for scanner.Scan() {
//     fmt.Println(scanner.Text())
// }
