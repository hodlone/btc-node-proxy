package listener

import (
	"btc-node-proxy/listener/gobHandler"
	"btc-node-proxy/msq"
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

		log.Printf("Topic: %v", topic)
		switch topic {
		case "hashblock":
			msq.Qpub(msq.HashBlock, body[0])
		case "rawblock":
			msq.Qpub(msq.RawBlock, body[0])
		case "hashtx":
			msq.Qpub(msq.HashTx, body[0])
		case "rawtx":
			gobHandler.RawTx(body[0])
			msq.Qpub(msq.RawTx, body[0])
		}
	}
}
