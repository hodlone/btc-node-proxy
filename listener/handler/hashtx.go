package handler

import (
	"btc-node-proxy/msq"
	"log"
	"strconv"
)

var hashTxCounter = 0

// HashTx handles the zmq messages published to the "hashtx" socket topic.
func HashTx(msg []byte) {
	hashTxCounter++
	log.Printf("Publishing hash tx %v", hashTxCounter)
	msq.Publish("btc.node.zmq.hashtx", []byte(strconv.Itoa(hashTxCounter)))

	// Print Transaction Hash
	// hash := hex.EncodeToString(msg)
	// log.Printf("From HashTx: %v", hash)
}
