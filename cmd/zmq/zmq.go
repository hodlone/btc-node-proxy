package zmq

import (
	"os"

	"github.com/NodeHodl/btc-node-proxy/pkg/logger"

	"github.com/pebbe/zmq4"
)

var (
	btcNodeZmqAddr = os.Getenv("BTC_NODE_ZMQ_ADDR")
)

var zmqLogger = logger.New("zmq")

// Start ...
func Start() {
	// Setup subscriber
	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	subscriber.SetLinger(0)
	subscriber.Connect(btcNodeZmqAddr)

	// Add topics
	subscriber.SetSubscribe("hashblock")
	subscriber.SetSubscribe("hashtx")
	subscriber.SetSubscribe("rawblock")
	subscriber.SetSubscribe("rawtx")

	zmqLogger.Infof("ZMQ Listening on %v", btcNodeZmqAddr)

	for {
		frames, err := subscriber.RecvMessageBytes(0)
		if err != nil {
			zmqLogger.Errorln(err)
			continue
		}

		topic := string(frames[0])
		// body := frames[1:]
		switch topic {
		case "hashblock":
			zmqLogger.Infoln("hashblock")
			// HashBlock(body[0])
		case "rawblock":
			zmqLogger.Infoln("rawblock")
			// RawBlock(body[0])
		case "hashtx":
			zmqLogger.Infoln("hashtx")
			// HashTx(body[0])
		case "rawtx":
			zmqLogger.Infoln("rawtx")
			// RawTx(body[0])
		}
	}
}
