package main

import (
	"bfio"
	"fmt"
	"net"
)

func msin() {
	ln, err := net.Listen('tcp', '8080')
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer ln.close()

	fmt.Println("Server is running on port 8080...")
}