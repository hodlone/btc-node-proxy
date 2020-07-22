package handler

import (
	"encoding/hex"
	"log"
)

func HashBlock(msg []byte) {
	hash := hex.EncodeToString(msg)
	log.Printf("From HashBlock: %v", hash)
}
