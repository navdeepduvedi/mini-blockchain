package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var receivedMessage []string
var sender string
var transactions = make(map[string]int)

func startClient1(address string, uid int) {

	sender = strconv.Itoa(uid)
	var fileAddress = "E:/transactions/" + sender

	//creating file to store transactions
	var f, _ = os.Create(fileAddress)

	//connect to this socket
	connClient, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
	}

	for {
		go startListening(connClient, f)
		if err != nil {
			fmt.Println(err)
		}
		//read in input from stdin
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter amount to transact : ")

		amount, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else if strings.TrimRight(amount, "\r\n") == "close" {
			//passing close command to server

			transaction := "close"
			connClient.Write([]byte(transaction + "\n"))
			break
		}
		var txnId = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())

		amount = strings.TrimRight(amount, "\r\n")

		//creating transaction string
		transaction := sender + "," + amount + "," + txnId + "," + "false"
		// send to socket
		// fmt.Fprint(connClient, text+"\n")

		connClient.Write([]byte(transaction + "\n"))

	}

}

func startListening(connClient net.Conn, fileName *os.File) {
	for {
		// reading messages from server
		message, err := bufio.NewReader(connClient).ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %+v", err.Error())
			return
		}
		receivedMessage = append(receivedMessage, message)
		data := strings.Split(message, ",")

		//writing transcation to file
		go writeToFile(receivedMessage, data, fileName)

		a, _ := strconv.Atoi(sender)
		b, _ := strconv.Atoi(data[0])

		if strings.TrimRight(data[3], "\r\n") == "false" && a != b {

			//returning acknowledgement for transaction received
			go acknowledgeTransaction(connClient, message, data)
		}
	}
}

func acknowledgeTransaction(connClient net.Conn, message string, data []string) {

	acknowledgeMessage := data[0] + "," + data[1] + "," + data[2] + "," + "true"
	_, err := connClient.Write([]byte(acknowledgeMessage + "\n"))
	if err != nil {
		fmt.Println(err)
	}
}
func writeToFile(receivedMessage []string, data []string, fileName *os.File) {
	if _, ok := transactions[data[2]]; ok {
		transactions[data[2]] += 1
	} else {
		transactions[data[2]] += 1
	}

	if transactions[data[2]] >= 2 && transactions[data[2]] < 3 {
		fileName.WriteString(data[0] + "," + data[1] + "," + data[2] + "\n")
	}
}
