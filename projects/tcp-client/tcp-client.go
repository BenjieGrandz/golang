package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // if server is listening a tcp connection is established
	if err != nil {
		log.Println("Connection error: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type messages and press < ENTER > .")

	for {
		fmt.Print("> ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if text == "exit\n" {
			fmt.Println("Closing connection.")
			break
		}

		conn.Write([]byte(text))

		reply, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Server replied: ", reply)
	}
}
