package main

import (
	"bufio"
	"fmt"
	"net"
)

func msin() {
	ln, err := net.Listen('tcp', '8080')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer ln.Close()
	fmt.Println("Server is running on port 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accepting error: ", err)
			continue
		}
	}
}