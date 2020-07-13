package main

import (
	"log"
	"os"
	"sync"

	"btc-node-proxy/listener"
	"btc-node-proxy/server"
	"btc-node-proxy/nats"
)

var (
	btcNodeZmqAddr = os.Getenv("BTC_NODE_ZMQ_ADDR")
	serverPort     = os.Getenv("PORT")
	natsAddr 			 = os.Getenv("NATS_ADDR")
	natsName 			 = os.Getenv("NATS_NAME")
)

func main() {

	log.Printf("env var %v\n", btcNodeZmqAddr)
	wg := new(sync.WaitGroup)

	wg.Add(3)

	go func() {
		nats.Start(natsAddr, natsName)
		wg.Done()
	}

	go func() {
		server.Start(serverPort)
		wg.Done()
	}()

	go func() {
		listener.Start(btcNodeZmqAddr)
		wg.Done()
	}()

	wg.Wait()
}
