package main

import (
	"os"
	"sync"

	"btc-node-proxy/listener"
	"btc-node-proxy/msq"
	"btc-node-proxy/server"
	"btc-node-proxy/vault"
)

var (
	btcNodeZmqAddr = os.Getenv("BTC_NODE_ZMQ_ADDR")
	serverPort     = os.Getenv("PORT")
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(4)

	go func() {
		msq.StartStanClient()
		wg.Done()
	}()

	go func() {
		server.Start(serverPort)
		wg.Done()
	}()

	go func() {
		listener.Start(btcNodeZmqAddr)
		wg.Done()
	}()

	go func() {
		vault.Start()
		wg.Done()
	}()

	wg.Wait()
}
