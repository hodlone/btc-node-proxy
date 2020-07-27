package msq

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var (
	NatsAddr = os.Getenv("NATS_ADDR")
	NatsName = os.Getenv("NATS_NAME")
)

const (
	HashBlock = "btc.node.zmq.hashblock"
	HashTx    = "btc.node.zmq.hashtx"
	RawBlock  = "btc.node.zmq.rawblock"
	RawTx     = "btc.node.zmq.rawtx"
	GobRawTx  = "btc.node.zmq.gobrawtx"
)

func Qpub(s string, m []byte) {
	NATS, err := nats.Connect(NatsAddr, nats.Name(NatsName))
	if err != nil {
		log.Fatal(err)
	}
	// defer connection close
	defer NATS.Close()

	NATS.Publish(s, m)

	NATS.Flush()
	if err := NATS.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s]\n", s)
	}
}
