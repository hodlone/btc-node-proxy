package main

import (
	"log"
	"sync"

	"btc-node-proxy/listener"
	"btc-node-proxy/server"
)

var (
	btcNodeZmqAddr = "tcp://bitcoin-core.bitcoin-test.svc.cluster.local:29000"
	serverPort     = "4500"
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
