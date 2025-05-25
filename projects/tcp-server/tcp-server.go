package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer ln.Close()
	fmt.Println("Server is running on port 8080...")

	for {
		// accepts, waits for and returns connection to listener object
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accepting error: ", err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Print("Received: ", msg)

		conn.Write([]byte("Echo: "+ msg))
	}
}