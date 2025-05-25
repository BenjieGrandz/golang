package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}

	defer listener.Close()
	log.Println("Server is running on http://localhost:8080")

	for {
		// accepts, waits for and returns connection to listener object
		conn, err := listener.Accept() // returns <net.Conn object> a 2 way connection
		if err != nil {
			fmt.Println("Connection error: ", err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected: ", err)
			return
		}

		fmt.Print("Received: ", message)

		conn.Write([]byte("Echo: " + message))
	}
}
