package handler

import (
	"encoding/json"
	"log"

	"github.com/piotrnar/gocoin/lib/btc"
)

func RawBlock(msg []byte) {
	bl, _ := btc.NewBlock(msg)
	bl.BuildTxList()
	json, err := json.Marshal(bl)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("From RawBlock Hash: %v", string(json))
	queue.Qpub("btc.rawblock", json)
}
