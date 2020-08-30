package handler

import (
	"btc-node-proxy/msq"
	"log"
)

// HashBlock handles the zmq messages published to the "hashblock" socket topic.
func HashBlock(msg []byte) {
	log.Println("Publishing hash block")
	msq.Publish("btc.node.zmq.hashblock", msg)

	// Print Blocks Hash
	// hash := hex.EncodeToString(msg)
	// log.Printf("From HashBlock: %v", hash)
}
