package handler

import (
	"btc-node-proxy/msq"
)

// HashBlock handles the zmq messages published to the "hashblock" socket topic.
func HashBlock(msg []byte) {
	msq.Qpub("btc.node.zmq.hashblock", msg)

	// Print Blocks Hash
	// hash := hex.EncodeToString(msg)
	// log.Printf("From HashBlock: %v", hash)
}
