package main

import (
	"os"
	"sync"

	"btc-node-proxy/listener"
	"btc-node-proxy/qeue"
	"btc-node-proxy/server"
)

var (
	btcNodeZmqAddr = os.Getenv("BTC_NODE_ZMQ_ADDR")
	serverPort     = os.Getenv("PORT")
	natsAddr       = os.Getenv("NATS_ADDR")
	natsName       = os.Getenv("NATS_NAME")
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(3)

	go func() {
		qeue.Start(natsAddr, natsName)
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
