package listener

import (
	"encoding/hex"
	"log"

	"github.com/pebbe/zmq4"
	// "btc-node-proxy/qeue"
)

func decode(src []byte) {
	dst := make([]byte, hex.DecodedLen(len(src)))
	log.Printf("dst: %v", dst)
	log.Printf("HEX: %v", hex.EncodeToString(src))

}

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
		e, addr, v, err := subscriber.RecvEvent(0)
		if err != nil {
			log.Println(err)
			continue
		}

		// utiliza la funcion qeue.Qpub(subject string, message []byte) para mandar los mensajes para nats.

		// los subjects disponibles son

		// "btc.rawtx"
		// "btc.rawblock"
		// "btc.hashtx"
		// "btc.hashblock"

		// log
		log.Printf("Event: %v", e)
		addrs := []byte(addr)
		log.Printf("Address: %v", addr)
		decode(addrs)
		log.Printf("Value: %v", v)

	}
}
