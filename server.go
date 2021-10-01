package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var listOfClients []*net.TCPConn

func startServer(address string) {
	fmt.Println("Starting...")

	//listen on all interfaces
	server, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
	}
	defer server.Close()
	for {
		// accept connection on port
		connServer, err := server.Accept()
		fmt.Println("accepting...")
		tcp := (connServer.(*net.TCPConn))
		if err != nil {
			log.Println("Error accepting request:", err)
		}

		listOfClients = append(listOfClients, tcp)
		fmt.Println(listOfClients)
		go handleRequest(connServer)
	}

}

func handleRequest(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		// iterating over the listOfClients
		fmt.Println(listOfClients)
		for i := range listOfClients {
			// sending message to all clients in the socket
			_, err := listOfClients[i].Write([]byte(message))
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}

// func broadCast(message string) {
// 	print("broadcasting")
// 	// iterating over the listOfClients

// 	for i := range listOfClients {
// 		// sending message to all clients in the socket
// 		_, err := listOfClients[i].Write([]byte(message))
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 	}
// 	print("broadcasting")
// }
