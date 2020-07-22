package handler

import (
	"encoding/json"
	"log"

	"btc-node-proxy/qeue"

	"github.com/piotrnar/gocoin/lib/btc"
)

// RawBlock, handles the zmq messages published to the "rawblock" socket topic.
func RawBlock(msg []byte) {
	bl, _ := btc.NewBlock(msg)
	bl.BuildTxList()

	// printTxs(bl)
	json, err := json.Marshal(bl)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("From RawBlock Hash: %v", string(json))
	qeue.Qpub("btc.rawblock", json)
}

func printTxs(bl *btc.Block) {
	for _, tx := range bl.Txs {
		log.Printf("Tx Hash: %v", tx.Hash)
		for _, output := range tx.TxOut {
			s := output.String(true)
			log.Printf("Tx Out Value: %v", s)
		}
	}
}
