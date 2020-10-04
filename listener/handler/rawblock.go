package handler

import (
	"github.com/NodeHodl/btc-node-proxy/msq"
)

// RawBlock handles the zmq messages published to the "rawblock" socket topic.
func RawBlock(msg []byte) {
	msq.Publish("btc.node.zmq.rawblock", msg)
}

// func RawBlock(msg []byte) {
// 	qeue.Qpub("btc.rawblock", msg)
//
// 	//RawBlock deserialize the block and return it.
// 	var msgBlock wire.MsgBlock
// 	err := msgBlock.Deserialize(bytes.NewReader(msg))
//
// 	printBlock(msgBlock)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
//
// func printBlock(msgBlock wire.MsgBlock) {
// 	log.Printf("#############")
// 	log.Printf("### BLOCK ###")
// 	log.Println("#############")
//
// 	log.Printf("Hash:                  %v", msgBlock.BlockHash())
// 	log.Printf("Block Version:         %v", msgBlock.Header.Version)
// 	log.Printf("Prev Block Hash:       %v", msgBlock.Header.PrevBlock)
// 	log.Printf("Merkle Tree Root Hash: %v", msgBlock.Header.MerkleRoot)
// 	log.Printf("Block Created:         %v", msgBlock.Header.Timestamp)
// 	log.Printf("Difficulty target:     %v", msgBlock.Header.Bits)
// 	log.Printf("Nonce used:            %v", msgBlock.Header.Nonce)
//
// 	for i, tx := range msgBlock.Transactions {
// 		log.Printf("### TX%v ###", i)
// 		log.Printf("Tx Version:            %v", tx.Version)
// 		log.Printf("Tx LockTime:           %v", tx.LockTime)
//
// 		for j, out := range tx.TxOut {
// 			log.Printf("### TX%v OUT%v ###", i, j)
// 			log.Printf("Tx Value BTC:          %v", out.Value)
// 		}
//
// 	}
// }
