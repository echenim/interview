package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listen.Close()
	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Connected to %s\n", conn.RemoteAddr().String())

	welcomeMsg := "Welcome to the TCP server!\n"
	conn.Write([]byte(welcomeMsg))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("Received from client: %s\n", text)
		conn.Write([]byte("Echo: " + text + "\n"))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
