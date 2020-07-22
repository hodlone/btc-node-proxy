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

	log.Println("Listening")
	for {
		frames, err := subscriber.RecvMessageBytes(0)
		if err != nil {
			log.Println(err)
			continue
		}

		topic := string(frames[0])
		body := frames[1:]

		log.Printf("topic: %v", topic)
		switch topic {
		case "hashblock":
			handler.HashBlock(body[0])
		case "rawblock":
			handler.RawBlock(body[0])
			// case "hashtx":
			// hexi := hex.EncodeToString(body[0])
			// log.Printf("From HashTx HEX: %v", hexi)
			// handler.HashTx(topic)
			// case "rawtx":
			// tx, _ := btc.NewTxIn(body[0])
			// log.Printf("From RawTx Hash: %v", tx.Sequence)
			// handler.RawTx(topic)
		}
	}
}
