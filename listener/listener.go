package listener

import (
	handler "btc-node-proxy/listener/handler"
	"log"

	"github.com/pebbe/zmq4"
)

// Start ...
func Start(btcNodeZmqAddr string) {

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
		body := frames[1:]
		switch topic {
		case "hashblock":
			handler.HashBlock(body[0])
		case "rawblock":
			handler.RawBlock(body[0])
		case "hashtx":
			handler.HashTx(body[0])
		case "rawtx":
			handler.RawTx(body[0])
		}
	}
}
