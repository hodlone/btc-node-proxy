package gobHandler

import (
	"btc-node-proxy/msq"
	"bytes"
	"log"

	"github.com/btcsuite/btcd/wire"
	"github.com/nats-io/nats.go"
)

// RawTx handles the zmq messages published to the "rawtx" socket topic.
func RawTx(msg []byte) {
	msgTx := newRawTx(msg)
	publishRawTx(msgTx)
}

func newRawTx(msg []byte) wire.MsgTx {
	// 	//RawBlock deserialize the block and return it.
	var msgTx wire.MsgTx

	err := msgTx.Deserialize(bytes.NewReader(msg))

	if err != nil {
		log.Fatal(err)
	}

	return msgTx
}

func publishRawTx(tx wire.MsgTx) {
	qpub(msq.GobRawTx, tx)
}

func qpub(s string, tx wire.MsgTx) {
	NATS, err := nats.Connect(msq.NatsAddr, nats.Name(msq.NatsName))
	if err != nil {
		log.Fatal(err)
	}
	// defer connection close
	defer NATS.Close()
	ec, err := nats.NewEncodedConn(NATS, nats.GOB_ENCODER)

	ec.Publish(s, tx)

	ec.Flush()
	if err := ec.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s]\n", s)
	}
}
