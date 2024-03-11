package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Failed to listen on port 8080: %v\n", err)
		os.Exit(1)
	}
	defer listen.Close()
	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Connected to %s\n", conn.RemoteAddr().String())

	welcomeMsg := "Welcome to the TCP server! Type 'exit' to disconnect.\n"
	_, err := conn.Write([]byte(welcomeMsg))
	if err != nil {
		fmt.Printf("Error sending welcome message: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			fmt.Printf("Client %s requested to close the connection.\n", conn.RemoteAddr().String())
			return
		}
		fmt.Printf("Received from client: %s\n", text)
		_, err := conn.Write([]byte("Echo: " + text + "\n"))
		if err != nil {
			fmt.Printf("Error sending echo message: %v\n", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Printf("Error reading from connection: %v\n", err)
		} else {
			fmt.Printf("Client %s disconnected.\n", conn.RemoteAddr().String())
		}
	}
}
