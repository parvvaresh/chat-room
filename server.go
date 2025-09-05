package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Store connected clients with their address
var clients = make(map[net.Conn]string)

// Handle each client connection
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Add client to the map
	clients[conn] = conn.RemoteAddr().String()
	fmt.Fprintf(conn, "Welcome! Your address: %s\n", conn.RemoteAddr().String())

	// Read messages from this client
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("[%s]: %s\n", clients[conn], message)
		// Send message to all other clients
		broadcast(conn, message)
	}

	// Remove client when disconnected
	delete(clients, conn)
}

// Broadcast a message to all clients except the sender
func broadcast(sender net.Conn, message string) {
	for client := range clients {
		if client != sender {
			fmt.Fprintf(client, "[%s]: %s\n", clients[sender], message)
		}
	}
}

func main() {
	// Start TCP server on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is running on port 8080...")

	// Accept new client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// Handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}
