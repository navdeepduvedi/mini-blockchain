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
	//connect to this socket
	connClient, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
	}
	sender = strconv.Itoa(uid)

	for {
		go startListening(connClient)
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
			fmt.Println(receivedMessage)
			connClient.Close()
			break

		}
		var txnId = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())

		amount = strings.TrimRight(amount, "\r\n")

		transaction := sender + "," + amount + "," + txnId + "," + "false"
		// send to socket
		// fmt.Fprint(connClient, text+"\n")

		n, _ := connClient.Write([]byte(transaction + "\n"))
		fmt.Println(n)

	}

}

func startListening(connClient net.Conn) {
	for {
		// reading messages from server
		message, err := bufio.NewReader(connClient).ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %+v", err.Error())
			return
		}
		receivedMessage = append(receivedMessage, message)
		data := strings.Split(message, ",")
		go writeToFile(receivedMessage, data)
		a, _ := strconv.Atoi(sender)
		b, _ := strconv.Atoi(data[0])
		if strings.TrimRight(data[3], "\r\n") == "false" && a != b {
			go acknowledgeTransaction(connClient, message, data)
		}
	}
}

func acknowledgeTransaction(connClient net.Conn, message string, data []string) {

	acknowledgeMessage := data[0] + "," + data[1] + "," + data[2] + "," + "true"
	n, err := connClient.Write([]byte(acknowledgeMessage + "\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
func writeToFile(receivedMessage []string, data []string) {
	f, _ := os.Create("E:/transactions/data")
	if _, ok := transactions[data[2]]; ok {
		transactions[data[2]] += 1
	} else {
		transactions[data[2]] += 1
	}

	if transactions[data[2]] >= 2 {
		result, _ := f.WriteString(data[0] + data[1] + data[2] + data[3])
		fmt.Println(result)
	}
}
