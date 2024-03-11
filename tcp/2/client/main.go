package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to the server. Type 'exit' to disconnect.")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			fmt.Println("Disconnecting from server.")
			break
		}

		_, err := fmt.Fprintf(conn, text+"\n")
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			break
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			break
		}
		fmt.Printf("Received from server: %s", response)
	}
}
