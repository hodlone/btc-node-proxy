package main

import (
	"sync"

	"github.com/NodeHodl/btc-node-proxy/bitcoin"
	"github.com/NodeHodl/btc-node-proxy/msq"
)

func main() {
	// Clients are initialized
	bitcoin.InitRPCClient()
	msq.InitStanClient()

	wg := new(sync.WaitGroup)

	wg.Add(3)

	// Servers are started
	go func() {
		StartHealthServer()
		wg.Done()
	}()

	go func() {
		StartZmqListener()
		wg.Done()
	}()

	go func() {
		StartNetworkMonitor()
		wg.Done()
	}()

	wg.Wait()
}
