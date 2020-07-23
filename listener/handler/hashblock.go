package handler

import (
	"btc-node-proxy/qeue"
	"encoding/hex"
	"log"
)

// HashBlock handles the zmq messages published to the "hashblock" socket topic.
func HashBlock(msg []byte) {
	qeue.Qpub("btc.hashblock", msg)

	// Print Blocks Hash
	hash := hex.EncodeToString(msg)
	log.Printf("From HashBlock: %v", hash)
}
