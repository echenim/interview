package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to the server")

	// Send a message to the server
	message := "Hello, Server!"
	fmt.Fprintf(conn, message+"\n")

	// Read the server's response
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
	} else {
		fmt.Print("Received from server: " + response)
	}
}
