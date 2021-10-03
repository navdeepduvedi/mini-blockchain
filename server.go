package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

		//storing client address connecting to server
		listOfClients = append(listOfClients, tcp)
		go handleRequest(connServer)
	}

}

func handleRequest(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		// sample process for string received
		// iterating over the listOfClients

		for i := range listOfClients {
			// sending message to all clients in the socket
			_, err := listOfClients[i].Write([]byte(message))
			if err != nil {
				fmt.Println(err)
			}
			if strings.TrimSpace(string(message)) == "close" {
				//closing connection to client

				fmt.Println("ho")
				break
			}

		}
	}
}
