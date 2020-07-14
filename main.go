package main

import (
	"log"
	"os"
	"sync"

	"btc-node-proxy/listener"
	"btc-node-proxy/server"
)

var (
	btcNodeZmqAddr = os.Getenv("os.Getenv(key)")
	serverPort     = os.Getenv("PORT")
)

func main() {

	log.Printf("env var %v\n", btcNodeZmqAddr)
	wg := new(sync.WaitGroup)

	wg.Add(2)

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
