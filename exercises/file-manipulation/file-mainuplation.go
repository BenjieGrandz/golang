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

	for _, name := range names {
		_, err := file.WriteString(name + "\n")
		if err != nil {
			fmt.Println("Error writting name to file :", err)
			return 
		}
	}

	fmt.Println("Names have been written into the file")

	// read from the file

	file, err = os.Open("names.txt")
	if err != nil {
		fmt.Println("Can't read file due to the following error", err)
		return
	}

	fmt.Println("Reading the names in the file")

	
}

// file, err := os.Create("file.txt")
// defer file.Close()

// file.WriteString("Hello\n")

// file, _ := os.Open("file.txt")
// scanner := bufio.NewScanner(file)
// for scanner.Scan() {
//     fmt.Println(scanner.Text())
// }
