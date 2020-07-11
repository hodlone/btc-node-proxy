package listener

import (
	"log"

	"github.com/pebbe/zmq4"
)

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
		s, err := subscriber.Recv(0)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("rec", s)
	}
}
