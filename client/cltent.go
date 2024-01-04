package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	reader := bufio.NewReader(os.Stdin)
	for {
		// Read user input
		fmt.Print("Enter username: ")
		username, _ := reader.ReadString('\n')
		fmt.Print("Enter password: ")
		password, _ := reader.ReadString('\n')
		// Send the username to the server

		conn.Write([]byte(username))
		conn.Write([]byte(password))

		if username[:len(username)-2] == "std1" && password[:len(password)-2] == "p@ssw0rd" {
			fmt.Printf("Hello\n")
		} else {
			fmt.Printf("Invalid credentials\n")
		}
		if username[:len(username)-2] == ":quit" {
			return
		}

		// Receive and print the server's response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Printf("Server response: %s", buffer[:n])
	}
}
