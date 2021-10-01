package main

import (
	"flag"
	"math/rand"
	"time"
)

type Transaction struct {
	TxnId  string
	Amount string
	Flag   int
}

var uid = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	host := flag.String("host", "", "Pass the host to start server")
	flag.Parse()

	if *host != "" {
		startClient1(*host, uid.Int())
	} else {
		startServer("localhost:8085")
	}
}
