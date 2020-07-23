package handler

import (
	"btc-node-proxy/qeue"
	"encoding/hex"
	"log"
)

// HashTx handles the zmq messages published to the "hashtx" socket topic.
func HashTx(msg []byte) {
	qeue.Qpub("btc.hashtx", msg)

	// Print Transaction Hash
	hash := hex.EncodeToString(msg)
	log.Printf("From HashTx: %v", hash)
}
