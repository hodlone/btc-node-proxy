package main

import (
	"os"
	"sync"

	"btc-node-proxy/listener"
	"btc-node-proxy/msq"
	"btc-node-proxy/server"
)

var (
	btcNodeZmqAddr = os.Getenv("BTC_NODE_ZMQ_ADDR")
	serverPort     = os.Getenv("PORT")
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(3)

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

	wg.Wait()
}
