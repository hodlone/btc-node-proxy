package main

import (
	"log"
	"os"

	// handler "github.com/NodeHodl/btc-node-proxy/listener/handler"

	"github.com/pebbe/zmq4"
)

var (
	btcNodeZmqAddr = os.Getenv("BTC_NODE_ZMQ_ADDR")
)

// StartZmqListener ...
func StartZmqListener() {

	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	subscriber.SetLinger(0)

	subscriber.Connect(btcNodeZmqAddr)

	subscriber.SetSubscribe("hashblock")
	subscriber.SetSubscribe("hashtx")
	subscriber.SetSubscribe("rawblock")
	subscriber.SetSubscribe("rawtx")

	log.Println("ZMQ Listening")
	for {
		frames, err := subscriber.RecvMessageBytes(0)
		if err != nil {
			log.Println(err)
			continue
		}

		topic := string(frames[0])
		// body := frames[1:]
		switch topic {
		case "hashblock":
			// handler.HashBlock(body[0])
		case "rawblock":
			log.Println("GOT NEW BLOCK ON ZMQ port 29000!")
			// handler.RawBlock(body[0])
		case "hashtx":
			// handler.HashTx(body[0])
		case "rawtx":
			log.Println("GOT NEW TX ON ZMQ port 29000!")
			// handler.RawTx(body[0])
		}
	}
}
